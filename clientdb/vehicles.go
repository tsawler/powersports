package clientdb

import (
	"database/sql"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
)

// VehicleModel holds the db connection
type VehicleModel struct {
	DB *sql.DB
}

// GetVehiclesForSaleByType returns slice of vehicles by type
func (m *VehicleModel) GetVehiclesForSaleByType(vehicleType int) ([]clientmodels.Vehicle, error) {
	var v []clientmodels.Vehicle

	query := `
		select 
		       id, 
		       stock_no, 
		       coalesce(cost, 0),
		       vin, 
		       coalesce(odometer, 0),
		       coalesce(year, 0),
		       coalesce(trim, ''),
		       vehicle_type,
		       coalesce(body, ''),
		       coalesce(seating_capacity,''),
		       coalesce(drive_train,''),
		       coalesce(engine,''),
		       coalesce(exterior_color,''),
		       coalesce(interior_color,''),
		       coalesce(transmission,''),
		       coalesce(options,''),
		       coalesce(model_number, ''),
		       coalesce(total_msr,0.0),
		       v.status,
		       coalesce(description, ''),
		       vehicle_makes_id,
		       vehicle_models_id,
		       hand_picked,
		       used,
		       coalesce(price_for_display,''),
		       created_at,
		       updated_at
		from 
		     wheelsanddeals.vehicles v 
		where
			vehicle_type = ?
			and status = 1

		order by year desc`

	rows, err := m.DB.Query(query, vehicleType)

	if err != nil {
		fmt.Println(err)
		return v, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &clientmodels.Vehicle{}
		err = rows.Scan(
			&c.ID,
			&c.StockNo,
			&c.Cost,
			&c.Vin,
			&c.Odometer,
			&c.Year,
			&c.Trim,
			&c.VehicleType,
			&c.Body,
			&c.SeatingCapacity,
			&c.DriveTrain,
			&c.Engine,
			&c.ExteriorColour,
			&c.InteriorColour,
			&c.Transmission,
			&c.Options,
			&c.ModelNumber,
			&c.TotalMSR,
			&c.Status,
			&c.Description,
			&c.VehicleMakesID,
			&c.VehicleModelsID,
			&c.HandPicked,
			&c.Used,
			&c.PriceForDisplay,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return v, err
		}

		// get make
		vehicleMake := clientmodels.Make{}

		query = `
			SELECT 
				id, 
				make, 
				created_at, 
				updated_at 
			FROM 
				wheelsanddeals.vehicle_makes 
			WHERE 
				id = ?`
		makeRow := m.DB.QueryRow(query, c.VehicleMakesID)

		err = makeRow.Scan(
			&vehicleMake.ID,
			&vehicleMake.Make,
			&vehicleMake.CreatedAt,
			&vehicleMake.UpdatedAt)
		if err != nil {
			fmt.Println("*** Error getting make:", err)
			//return v, err
		}
		c.Make = vehicleMake

		// get model
		model := clientmodels.Model{}

		query = `
			SELECT 
				id, 
				model, 
				vehicle_makes_id,
				created_at, 
				updated_at 
			FROM 
				wheelsanddeals.vehicle_models 
			WHERE 
				id = ?`
		modelRow := m.DB.QueryRow(query, c.VehicleModelsID)

		err = modelRow.Scan(
			&model.ID,
			&model.Model,
			&model.MakeID,
			&model.CreatedAt,
			&model.UpdatedAt)
		if err != nil {
			fmt.Println("*** Error getting model:", err)
			//return v, err
		}
		c.Model = model

		// get options
		query = `
			select 
				vo.id, 
				vo.vehicle_id,
				vo.option_id,
				vo.created_at,
				vo.updated_at,
				o.option_name
			from 
				wheelsanddeals.vehicle_options vo
				left join wheelsanddeals.options o on (vo.option_id = o.id)
			where
				vo.vehicle_id = ?
				and o.active = 1
			order by 
				o.option_name`
		oRows, err := m.DB.Query(query, c.ID)
		if err != nil {
			fmt.Println("*** Error getting options:", err)
		}

		var vehicleOptions []*clientmodels.VehicleOption
		for oRows.Next() {
			o := &clientmodels.VehicleOption{}
			err = oRows.Scan(
				&o.ID,
				&o.VehicleID,
				&o.OptionID,
				&o.CreatedAt,
				&o.UpdatedAt,
				&o.OptionName,
			)

			if err != nil {
				fmt.Println(err)
			} else {
				vehicleOptions = append(vehicleOptions, o)
			}
		}
		c.VehicleOptions = vehicleOptions
		oRows.Close()

		// get images
		query = `
			select 
				id, 
				vehicle_id,
				image,
				created_at,
				updated_at,
				sort_order
			from 
				wheelsanddeals.vehicle_images 
			where
				vehicle_id = ?
			order by 
				sort_order`
		iRows, err := m.DB.Query(query, c.ID)
		if err != nil {
			fmt.Println(err)
		}

		var vehicleImages []*clientmodels.Image
		for iRows.Next() {
			o := &clientmodels.Image{}
			err = iRows.Scan(
				&o.ID,
				&o.VehicleID,
				&o.Image,
				&o.CreatedAt,
				&o.UpdatedAt,
				&o.SortOrder,
			)

			if err != nil {
				fmt.Println(err)
			} else {
				vehicleImages = append(vehicleImages, o)
			}
		}
		c.Images = vehicleImages
		iRows.Close()

		current := *c
		v = append(v, current)
	}

	return v, nil
}

// AllVehiclesPaginated returns paginated slice of vehicles, by type
func (m *VehicleModel) AllVehiclesPaginated(vehicleTypeID, perPage, offset, year, make, model, price int) ([]clientmodels.Vehicle, int, error) {
	var v []clientmodels.Vehicle

	where := ""
	orderBy := "order by year desc"
	if year > 0 {
		where = fmt.Sprintf("and v.year = %d", year)
	}

	if make > 0 {
		where = fmt.Sprintf("%s and v.vehicle_makes_id = %d", where, make)
	}

	if model > 0 {
		where = fmt.Sprintf("%s and v.vehicle_models_id = %d", where, model)
	}

	if price == 1 {
		orderBy = "order by v.cost asc"
	} else if price == 2 {
		orderBy = "order by v.cost desc"
	}

	stmt := ""
	var nRows *sql.Rows

	if vehicleTypeID < 1000 {
		stmt = fmt.Sprintf(`
		select 
			count(v.id) 
		from 
			wheelsanddeals.vehicles v 
		where 
			status = 1 
			and vehicle_type = ? %s`, where)
		n, err := m.DB.Query(stmt, vehicleTypeID)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		nRows = n
	} else if vehicleTypeID == 1000 {
		stmt = fmt.Sprintf(`
		select 
			count(v.id) 
		from 
			wheelsanddeals.vehicles v 
		where 
			status = 1 
			and vehicle_type in (8, 11, 12, 16, 13, 10, 7, 9, 15, 17, 14) %s 
			and v.used = 1`, where)
		n, err := m.DB.Query(stmt)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		nRows = n
	} else if vehicleTypeID == 1001 {
		stmt = fmt.Sprintf(`
		select 
			count(v.id) 
		from 
			wheelsanddeals.vehicles v 
		where 
			status = 1 
			and vehicle_type in (13, 10, 9, 15) %s 
			and v.used = 1`, where)
		n, err := m.DB.Query(stmt)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		nRows = n
	}

	defer nRows.Close()

	nRows.Next()
	var num int
	err := nRows.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}

	query := ""
	var rows *sql.Rows

	if vehicleTypeID < 1000 {
		query = fmt.Sprintf(`
		select 
		       id, 
		       stock_no, 
		       coalesce(cost, 0),
		       vin, 
		       coalesce(odometer, 0),
		       coalesce(year, 0),
		       coalesce(trim, ''),
		       vehicle_type,
		       coalesce(body, ''),
		       coalesce(seating_capacity,''),
		       coalesce(drive_train,''),
		       coalesce(engine,''),
		       coalesce(exterior_color,''),
		       coalesce(interior_color,''),
		       coalesce(transmission,''),
		       coalesce(options,''),
		       coalesce(model_number, ''),
		       coalesce(total_msr,0.0),
		       v.status,
		       coalesce(description, ''),
		       vehicle_makes_id,
		       vehicle_models_id,
		       hand_picked,
		       used,
		       coalesce(price_for_display,''),
		       created_at,
		       updated_at
		from 
		     wheelsanddeals.vehicles v 
		where
			vehicle_type = ?
			and status = 1
			%s
			%s
		limit ? offset ?`, where, orderBy)
		rows, err = m.DB.Query(query, vehicleTypeID, perPage, offset)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		defer rows.Close()
	} else if vehicleTypeID == 1000 {
		query = fmt.Sprintf(`
		select 
		       id, 
		       stock_no, 
		       coalesce(cost, 0),
		       vin, 
		       coalesce(odometer, 0),
		       coalesce(year, 0),
		       coalesce(trim, ''),
		       vehicle_type,
		       coalesce(body, ''),
		       coalesce(seating_capacity,''),
		       coalesce(drive_train,''),
		       coalesce(engine,''),
		       coalesce(exterior_color,''),
		       coalesce(interior_color,''),
		       coalesce(transmission,''),
		       coalesce(options,''),
		       coalesce(model_number, ''),
		       coalesce(total_msr,0.0),
		       v.status,
		       coalesce(description, ''),
		       vehicle_makes_id,
		       vehicle_models_id,
		       hand_picked,
		       used,
		       coalesce(price_for_display,''),
		       created_at,
		       updated_at
		from 
		     wheelsanddeals.vehicles v 
		where
			vehicle_type in (8, 11, 12, 16, 7, 17, 14)
			and status = 1
			and v.used = 1
			%s
			%s
		limit ? offset ?`, where, orderBy)

		rows, err = m.DB.Query(query, perPage, offset)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		defer rows.Close()
	} else if vehicleTypeID == 1001 {
		query = fmt.Sprintf(`
		select 
		       id, 
		       stock_no, 
		       coalesce(cost, 0),
		       vin, 
		       coalesce(odometer, 0),
		       coalesce(year, 0),
		       coalesce(trim, ''),
		       vehicle_type,
		       coalesce(body, ''),
		       coalesce(seating_capacity,''),
		       coalesce(drive_train,''),
		       coalesce(engine,''),
		       coalesce(exterior_color,''),
		       coalesce(interior_color,''),
		       coalesce(transmission,''),
		       coalesce(options,''),
		       coalesce(model_number, ''),
		       coalesce(total_msr,0.0),
		       v.status,
		       coalesce(description, ''),
		       vehicle_makes_id,
		       vehicle_models_id,
		       hand_picked,
		       used,
		       coalesce(price_for_display,''),
		       created_at,
		       updated_at
		from 
		     wheelsanddeals.vehicles v 
		where
			vehicle_type in (13, 10, 9, 15)
			and status = 1
			and v.used = 1
			%s
			%s
		limit ? offset ?`, where, orderBy)

		rows, err = m.DB.Query(query, perPage, offset)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		defer rows.Close()
	}

	for rows.Next() {
		c := &clientmodels.Vehicle{}
		err = rows.Scan(
			&c.ID,
			&c.StockNo,
			&c.Cost,
			&c.Vin,
			&c.Odometer,
			&c.Year,
			&c.Trim,
			&c.VehicleType,
			&c.Body,
			&c.SeatingCapacity,
			&c.DriveTrain,
			&c.Engine,
			&c.ExteriorColour,
			&c.InteriorColour,
			&c.Transmission,
			&c.Options,
			&c.ModelNumber,
			&c.TotalMSR,
			&c.Status,
			&c.Description,
			&c.VehicleMakesID,
			&c.VehicleModelsID,
			&c.HandPicked,
			&c.Used,
			&c.PriceForDisplay,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}

		// get make
		vehicleMake := clientmodels.Make{}

		query = `
			SELECT 
				id, 
				make, 
				created_at, 
				updated_at 
			FROM 
				wheelsanddeals.vehicle_makes 
			WHERE 
				id = ?`
		makeRow := m.DB.QueryRow(query, c.VehicleMakesID)

		err = makeRow.Scan(
			&vehicleMake.ID,
			&vehicleMake.Make,
			&vehicleMake.CreatedAt,
			&vehicleMake.UpdatedAt)
		if err != nil {
			fmt.Println("*** Error getting make:", err)
			//return v, err
		}
		c.Make = vehicleMake

		// get model
		model := clientmodels.Model{}

		query = `
			SELECT 
				id, 
				model, 
				vehicle_makes_id,
				created_at, 
				updated_at 
			FROM 
				wheelsanddeals.vehicle_models 
			WHERE 
				id = ?`
		modelRow := m.DB.QueryRow(query, c.VehicleModelsID)

		err = modelRow.Scan(
			&model.ID,
			&model.Model,
			&model.MakeID,
			&model.CreatedAt,
			&model.UpdatedAt)
		if err != nil {
			fmt.Println("*** Error getting model:", err)
			//return v, err
		}
		c.Model = model

		// get options
		query = `
			select 
				vo.id, 
				vo.vehicle_id,
				vo.option_id,
				vo.created_at,
				vo.updated_at,
				o.option_name
			from 
				wheelsanddeals.vehicle_options vo
				left join wheelsanddeals.options o on (vo.option_id = o.id)
			where
				vo.vehicle_id = ?
				and o.active = 1
			order by 
				o.option_name`
		oRows, err := m.DB.Query(query, c.ID)
		if err != nil {
			fmt.Println("*** Error getting options:", err)
		}

		var vehicleOptions []*clientmodels.VehicleOption
		for oRows.Next() {
			o := &clientmodels.VehicleOption{}
			err = oRows.Scan(
				&o.ID,
				&o.VehicleID,
				&o.OptionID,
				&o.CreatedAt,
				&o.UpdatedAt,
				&o.OptionName,
			)

			if err != nil {
				fmt.Println(err)
			} else {
				vehicleOptions = append(vehicleOptions, o)
			}
		}
		c.VehicleOptions = vehicleOptions
		oRows.Close()

		// get images
		query = `
			select 
				id, 
				vehicle_id,
				image,
				created_at,
				updated_at,
				sort_order
			from 
				wheelsanddeals.vehicle_images 
			where
				vehicle_id = ?
			order by 
				sort_order`
		iRows, err := m.DB.Query(query, c.ID)
		if err != nil {
			fmt.Println(err)
		}

		var vehicleImages []*clientmodels.Image
		for iRows.Next() {
			o := &clientmodels.Image{}
			err = iRows.Scan(
				&o.ID,
				&o.VehicleID,
				&o.Image,
				&o.CreatedAt,
				&o.UpdatedAt,
				&o.SortOrder,
			)

			if err != nil {
				fmt.Println(err)
			} else {
				vehicleImages = append(vehicleImages, o)
			}
		}
		c.Images = vehicleImages
		iRows.Close()

		current := *c
		v = append(v, current)
	}

	return v, num, nil
}

// GetYearsForVehicleType gets years for vehicle type
func (m *VehicleModel) GetYearsForVehicleType(id int) ([]int, error) {
	var years []int
	query := `
			select distinct 
				v.year
			from 
				wheelsanddeals.vehicles v
			where
				vehicle_type = ?
				and v.status = 1
			order by 
				year desc`
	rows, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var y int
		err = rows.Scan(&y)
		if err != nil {
			fmt.Println(err)
		}
		years = append(years, y)
	}
	return years, nil
}

// GetMakesForVehicleType gets makes for vehicle type
func (m *VehicleModel) GetMakesForVehicleType(id int) ([]clientmodels.Make, error) {
	var makes []clientmodels.Make
	query := ""

	if id < 1000 {
		query = `
			select  
				m.id, m.make
			from 
				wheelsanddeals.vehicle_makes m
			where
				m.id in (select v.vehicle_makes_id from wheelsanddeals.vehicles v where status = 1 and vehicle_type = ?)
			order by 
				m.make`
	} else if id == 1000 {
		query = `
			select  
				m.id, m.make
			from 
				wheelsanddeals.vehicle_makes m
			where
				m.id in (select v.vehicle_makes_id from wheelsanddeals.vehicles v where status = 1 and used = 1 and vehicle_type  in
				(8, 11, 12, 16, 7, 17, 14, ?)
)
			order by 
				m.make`
	} else if id == 1001 {
		query = `
			select  
				m.id, m.make
			from 
				wheelsanddeals.vehicle_makes m
			where
				m.id in (select v.vehicle_makes_id from wheelsanddeals.vehicles v where status = 1 and used = 1 and vehicle_type  in
				(13, 10, 9, 15, ?)
)
			order by 
				m.make`
	}
	rows, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var y clientmodels.Make
		err = rows.Scan(
			&y.ID,
			&y.Make)
		if err != nil {
			fmt.Println(err)
		}
		makes = append(makes, y)
	}

	return makes, nil
}

// GetModelsForVehicleType gets models for vehicle type
func (m *VehicleModel) GetModelsForVehicleType(id int) ([]clientmodels.Model, error) {
	var models []clientmodels.Model
	query := `
			select  
				m.id, m.model
			from 
				wheelsanddeals.vehicle_models m
			where
				m.id in (select v.vehicle_models_id from wheelsanddeals.vehicles v where status = 1 and vehicle_type = ?)
			order by 
				m.model`
	rows, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var y clientmodels.Model
		err = rows.Scan(
			&y.ID,
			&y.Model)
		if err != nil {
			fmt.Println(err)
		}
		models = append(models, y)
	}

	return models, nil
}

// GetPowerSportItem gets a complete record for a power sports item
func (m *VehicleModel) GetPowerSportItem(id int) (clientmodels.Vehicle, error) {
	var c clientmodels.Vehicle

	query := `
		select 
		       id, 
		       stock_no, 
		       coalesce(cost, 0),
		       vin, 
		       coalesce(odometer, 0),
		       coalesce(year, 0),
		       coalesce(trim, ''),
		       vehicle_type,
		       coalesce(body, ''),
		       coalesce(seating_capacity,''),
		       coalesce(drive_train,''),
		       coalesce(engine,''),
		       coalesce(exterior_color,''),
		       coalesce(interior_color,''),
		       coalesce(transmission,''),
		       coalesce(options,''),
		       coalesce(model_number, ''),
		       coalesce(total_msr,0.0),
		       v.status,
		       coalesce(description, ''),
		       vehicle_makes_id,
		       vehicle_models_id,
		       hand_picked,
		       used,
		       coalesce(price_for_display,''),
		       created_at,
		       updated_at
		from 
		     wheelsanddeals.vehicles v 
		where
			id = ?`

	row := m.DB.QueryRow(query, id)

	err := row.Scan(
		&c.ID,
		&c.StockNo,
		&c.Cost,
		&c.Vin,
		&c.Odometer,
		&c.Year,
		&c.Trim,
		&c.VehicleType,
		&c.Body,
		&c.SeatingCapacity,
		&c.DriveTrain,
		&c.Engine,
		&c.ExteriorColour,
		&c.InteriorColour,
		&c.Transmission,
		&c.Options,
		&c.ModelNumber,
		&c.TotalMSR,
		&c.Status,
		&c.Description,
		&c.VehicleMakesID,
		&c.VehicleModelsID,
		&c.HandPicked,
		&c.Used,
		&c.PriceForDisplay,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return c, err
	}

	// get make
	vehicleMake := clientmodels.Make{}

	query = `
			SELECT 
				id, 
				make, 
				created_at, 
				updated_at 
			FROM 
				wheelsanddeals.vehicle_makes 
			WHERE 
				id = ?`
	makeRow := m.DB.QueryRow(query, c.VehicleMakesID)

	err = makeRow.Scan(
		&vehicleMake.ID,
		&vehicleMake.Make,
		&vehicleMake.CreatedAt,
		&vehicleMake.UpdatedAt)
	if err != nil {
		fmt.Println("*** Error getting make:", err)
	}
	c.Make = vehicleMake

	// get model
	model := clientmodels.Model{}

	query = `
			SELECT 
				id, 
				model, 
				vehicle_makes_id,
				created_at, 
				updated_at 
			FROM 
				wheelsanddeals.vehicle_models 
			WHERE 
				id = ?`
	modelRow := m.DB.QueryRow(query, c.VehicleModelsID)

	err = modelRow.Scan(
		&model.ID,
		&model.Model,
		&model.MakeID,
		&model.CreatedAt,
		&model.UpdatedAt)
	if err != nil {
		fmt.Println("*** Error getting model:", err)
	}
	c.Model = model

	// get options
	query = `
			select 
				vo.id, 
				vo.vehicle_id,
				vo.option_id,
				vo.created_at,
				vo.updated_at,
				o.option_name
			from 
				wheelsanddeals.vehicle_options vo
				left join wheelsanddeals.options o on (vo.option_id = o.id)
			where
				vo.vehicle_id = ?
				and o.active = 1
			order by 
				o.option_name`
	oRows, err := m.DB.Query(query, c.ID)
	if err != nil {
		fmt.Println("*** Error getting options:", err)
	}

	var vehicleOptions []*clientmodels.VehicleOption
	for oRows.Next() {
		o := &clientmodels.VehicleOption{}
		err = oRows.Scan(
			&o.ID,
			&o.VehicleID,
			&o.OptionID,
			&o.CreatedAt,
			&o.UpdatedAt,
			&o.OptionName,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			vehicleOptions = append(vehicleOptions, o)
		}
	}

	c.VehicleOptions = vehicleOptions

	oRows.Close()

	// get images
	query = `
			select 
				id, 
				vehicle_id,
				image,
				created_at,
				updated_at,
				sort_order
			from 
				wheelsanddeals.vehicle_images 
			where
				vehicle_id = ?
			order by 
				sort_order`
	iRows, err := m.DB.Query(query, c.ID)
	if err != nil {
		fmt.Println(err)
	}

	defer iRows.Close()

	var vehicleImages []*clientmodels.Image
	for iRows.Next() {
		o := &clientmodels.Image{}
		err = iRows.Scan(
			&o.ID,
			&o.VehicleID,
			&o.Image,
			&o.CreatedAt,
			&o.UpdatedAt,
			&o.SortOrder,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			vehicleImages = append(vehicleImages, o)
		}
	}
	c.Images = vehicleImages
	iRows.Close()

	// get video, if any
	query = `
			select 
				v.id, 
				v.video_name,
				v.file_name,
				v.thumb
			from 
				wheelsanddeals.vehicle_videos vv
				left join wheelsanddeals.videos v on (vv.video_id = v.id)
			where
				vv.vehicle_id = ?
			limit 1`
	vRow := m.DB.QueryRow(query, c.ID)

	var vehicleVideo clientmodels.Video

	err = vRow.Scan(
		&vehicleVideo.ID,
		&vehicleVideo.VideoName,
		&vehicleVideo.FileName,
		&vehicleVideo.Thumb,
	)

	if err == nil {
		c.Video = vehicleVideo
	}

	c.Images = vehicleImages
	return c, nil
}

// GetSales gets max six sales people
func (m *VehicleModel) GetSales() ([]clientmodels.SalesStaff, error) {
	var s []clientmodels.SalesStaff

	query := `
		select 
		       id, 
		       coalesce(salesperson_name, ''),
		       coalesce(slug, ''),
		       coalesce(email, ''),
		       coalesce(phone, ''),
		       coalesce(image, '')
		from 
		     wheelsanddeals.sales 
		where
			active = 1


		order by RAND() limit 6`

	rows, err := m.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return s, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &clientmodels.SalesStaff{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.Email,
			&c.Phone,
			&c.Image,
		)
		if err != nil {
			fmt.Println(err)
			return s, err
		}
		staff := *c
		s = append(s, staff)
	}
	return s, nil
}

// InsertCreditApp saves a credit application
func (m *VehicleModel) InsertCreditApp(a clientmodels.CreditApp) error {
	stmt := `
	INSERT INTO wheelsanddeals.credit_applications (first_name, last_name, email, phone, address, city, province, zip, 
	                   vehicle, processed, created_at, updated_at)
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := m.DB.Exec(stmt,
		a.FirstName,
		a.LastName,
		a.Email,
		a.Phone,
		a.Address,
		a.City,
		a.Province,
		a.Zip,
		a.Vehicle,
		a.Processed,
		a.CreatedAt,
		a.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// InsertTestDrive saves a test drive application
func (m *VehicleModel) InsertTestDrive(a clientmodels.TestDrive) error {
	stmt := `
	INSERT INTO wheelsanddeals.test_drives (users_name, email, phone, preferred_date, preferred_time, vehicle_id, 
	                   processed, created_at, updated_at)
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := m.DB.Exec(stmt,
		a.UsersName,
		a.Email,
		a.Phone,
		a.PreferredDate,
		a.PreferredTime,
		a.VehicleID,
		a.Processed,
		a.CreatedAt,
		a.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// InsertQuickQuote saves a quick quote to remote dataabase
func (m *VehicleModel) InsertQuickQuote(a clientmodels.QuickQuote) error {
	stmt := `
	INSERT INTO wheelsanddeals.quick_quotes (users_name, email, phone, vehicle_id, 
	                   processed, created_at, updated_at)
    VALUES(?, ?, ?, ?, ?, ?, ?)
    `

	_, err := m.DB.Exec(stmt,
		a.UsersName,
		a.Email,
		a.Phone,
		a.VehicleID,
		a.Processed,
		a.CreatedAt,
		a.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
