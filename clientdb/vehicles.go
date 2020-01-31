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
		     vehicles v 
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
				vehicle_makes 
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
				vehicle_models 
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
				vehicle_options vo
				left join options o on (vo.option_id = o.id)
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
				vehicle_images 
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
			vehicles v 
		where 
			status = 1 
			and vehicle_type = ? %s`, where)
		n, err := m.DB.Query(stmt, vehicleTypeID)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		nRows = n
	} else {
		stmt = fmt.Sprintf(`
		select 
			count(v.id) 
		from 
			vehicles v 
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
		     vehicles v 
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
	} else {
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
		     vehicles v 
		where
			vehicle_type in (8, 11, 12, 16, 13, 10, 7, 9, 15, 17, 14)
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
				vehicle_makes 
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
				vehicle_models 
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
				vehicle_options vo
				left join options o on (vo.option_id = o.id)
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
				vehicle_images 
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
				vehicles v
			where
				vehicle_type = ?
				and v.status = 1
			order by 
				year desc`
	rows, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

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

	query := `
			select  
				m.id, m.make
			from 
				vehicle_makes m
			where
				m.id in (select v.vehicle_makes_id from vehicles v where status = 1 and vehicle_type = ?)
			order by 
				m.make`
	rows, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

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
				vehicle_models m
			where
				m.id in (select v.vehicle_models_id from vehicles v where status = 1 and vehicle_type = ?)
			order by 
				m.model`
	rows, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

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
