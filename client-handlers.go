package clienthandlers

import (
	"encoding/json"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	"github.com/tsawler/goblender/pkg/forms"
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/maildata"
	"github.com/tsawler/goblender/pkg/templates"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var vehicleModel *clientdb.VehicleModel

// JSONResponse is a generic struct to hold json responses
type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// CompareVehicles Show 2 or 3 vehicles in table TODO
func CompareVehicles(w http.ResponseWriter, r *http.Request) {
	ids := r.Form.Get("ids")
	infoLog.Println("Ids:", ids)
	helpers.Render(w, r, "compare.page.tmpl", &templates.TemplateData{})
}

// GetAllMotorcycles gets all motorcycles
func GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/motorcycle"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 7
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "motorcycle-inventory")
}

// GetAllBruteForce gets all brute force
func GetAllBruteForce(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-brute-force"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 8
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-brute-force")
}

// GetAllTeryx gets all teryx
func GetAllTeryx(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-teryx"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 12
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-teryx")
}

// GetAllMule gets all mules
func GetAllMule(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-mule"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 11
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-mule")
}

// GetAllJetSki gets all new jetskis
func GetAllJetSki(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/watercraft-jetski"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 13
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "watercraft-jetski-kawaski")
}

// GetAllMercuryOutboards gets all mercury outboards
func GetAllMercuryOutboards(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/outboard-motors-mercury"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 10
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "outboard-motors-mercury-outboards")
}

// GetAllElectricBikes gets all Pedegogo
func GetAllElectricBikes(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/electric-bikes"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 16
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "electric-bikes")
}

// GetAllScooters gets all scooters
func GetAllScooters(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 17
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "vespa-piaggio-gas-and-electric-scooters-mopeds")
}

// GetAllPontoonBoats gets all pontoon boats (bennington & crestliner
func GetAllPontoonBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/bennington-crestliner-pontoon-boats"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 9
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "pontoon-boats-bennington-and-crestliner")
}

// GetAllPowerBoats gets all power boats
func GetAllPowerBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/speed-boats-aluminum-boats-power-boats"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 15
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "speed-boats-aluminum-boats-power-boats")
}

// GetAllUsedPowerSports gets all used powersports
func GetAllUsedPowerSports(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/used-motorcycles-used-atv-used-boats-used-pontoons"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 1000
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "used-motorcycles-used-atv-used-boats-used-pontoons")
}

// GetAllTrailers gets all used powersports
func GetAllTrailers(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-trailers-boat-trailers-utility-trailers"
	stringMap["item-link-prefix"] = "view"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 14
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-trailers-boattrailers-utility-trailers")
}

// renderInventory renders inventory for a product type
func renderInventory(r *http.Request, stringMap map[string]string, vehicleType int, w http.ResponseWriter, intMap map[string]int, templateName, slug string) {
	var offset int
	var selectedYear, selectedMake, selectedModel, selectedPrice int
	pagerSuffix := ""
	stringMap["item-link-prefix"] = "view"
	stringMap["pager-prefix"] = "powersports-inventory"

	pageIndex, err := strconv.Atoi(r.URL.Query().Get(":pageIndex"))
	if err != nil {
		pageIndex = 1
	}

	searching, ok := r.URL.Query()["year"]
	if !ok || len(searching[0]) < 1 {
		selectedYear = 0
		selectedMake = 0
		selectedModel = 0
		selectedPrice = 0
	} else {
		selectedYear, _ = strconv.Atoi(r.URL.Query()["year"][0])
		selectedMake, _ = strconv.Atoi(r.URL.Query()["make"][0])
		selectedModel, _ = strconv.Atoi(r.URL.Query()["model"][0])
		selectedPrice, _ = strconv.Atoi(r.URL.Query()["price"][0])
		pagerSuffix = fmt.Sprintf("?year=%d&make=%d&model=%d&price=%d", selectedYear, selectedMake, selectedModel, selectedPrice)
		stringMap["pager-suffix"] = pagerSuffix
	}

	perPage := 10
	offset = (pageIndex - 1) * perPage

	vehicles, num, err := vehicleModel.AllVehiclesPaginated(vehicleType, perPage, offset, selectedYear, selectedMake, selectedModel, selectedPrice)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	rowSets := make(map[string]interface{})
	rowSets["vehicles"] = vehicles

	intMap["num-vehicles"] = num
	intMap["current-page"] = pageIndex
	intMap["year"] = selectedYear
	intMap["make"] = selectedMake
	intMap["model"] = selectedModel
	intMap["price"] = selectedPrice

	pg, err := pageModel.GetBySlug(slug)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// get makes
	makes, err := vehicleModel.GetMakesForVehicleType(vehicleType)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	// get models
	models, err := vehicleModel.GetModelsForVehicleType(vehicleType)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	// get years
	years, err := vehicleModel.GetYearsForVehicleType(vehicleType)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	rowSets["years"] = years
	rowSets["models"] = models
	rowSets["makes"] = makes

	helpers.Render(w, r, templateName, &templates.TemplateData{
		Page:      pg,
		RowSets:   rowSets,
		IntMap:    intMap,
		StringMap: stringMap,
	})
}

// ShowItem shows a product (i.e. ATV, whatever)
func ShowItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":ID"))

	pg, err := pageModel.GetBySlug("power-sports-item")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	pg.PageNotEditable = 1

	item, err := vehicleModel.GetPowerSportItem(id)
	if err != nil {
		fmt.Fprint(w, "custom 404")
		return
	}

	rowSets := make(map[string]interface{})
	rowSets["item"] = item

	staff, err := vehicleModel.GetSales()
	if err != nil {
		errorLog.Println(err)
	}

	rowSets["sales"] = staff

	helpers.Render(w, r, "item.page.tmpl", &templates.TemplateData{
		Page:    pg,
		RowSets: rowSets,
	})
}

// QuickQuote sends a quick quote request
func QuickQuote(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	interest := r.Form.Get("interested")

	content := fmt.Sprintf(`
		<p>
			<strong>PowerSports Quick Quote Request</strong>:<br><br>
			<strong>Name:</strong> %s <br>
			<strong>Email:</strong> %s <br>
			<strong>Phone:</strong> %s <br>
			<strong>Interested In:</strong><br><br>
			%s
		</p>
`, name, email, phone, interest)

	var cc []string
	cc = append(cc, "wheelsanddeals@pbssystems.com")

	mailMessage := maildata.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     "PowerSports Quick Quote Request",
		Content:     template.HTML(content),
		Template:    "generic-email.mail.tmpl",
		CC:          cc,
	}

	helpers.SendEmail(mailMessage)

	theData := JSONResponse{
		OK: true,
	}

	// build the json response from the struct
	out, err := json.MarshalIndent(theData, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// send json to client
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		errorLog.Println(err)
	}
}

// TestDrive sends a test drive request
func TestDrive(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	interest := r.Form.Get("interested")
	pDate := r.Form.Get("preferred_date")
	pTime := r.Form.Get("preferred_time")
	vid, _ := strconv.Atoi(r.Form.Get("vehicle_id"))

	content := fmt.Sprintf(`
		<p>
			<strong>PowerSports Test Drive Request</strong>:<br><br>
			<strong>Name:</strong> %s <br>
			<strong>Email:</strong> %s <br>
			<strong>Phone:</strong> %s <br>
			<strong>Preferred Date:</strong> %s<br>
			<strong>Preferred Time:</strong> %s<br>
			<strong>Interested In:</strong><br><br>
			%s
		</p>
`, name, email, phone, pDate, pTime, interest)

	var cc []string
	cc = append(cc, "wheelsanddeals@pbssystems.com")
	cc = append(cc, "john.eliakis@wheelsanddeals.ca")

	mailMessage := maildata.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     "PowerSports Test Drive Request",
		Content:     template.HTML(content),
		Template:    "generic-email.mail.tmpl",
		CC:          cc,
	}

	helpers.SendEmail(mailMessage)

	// save
	td := clientmodels.TestDrive{
		UsersName:     name,
		Email:         email,
		Phone:         phone,
		PreferredDate: pDate,
		PreferredTime: pTime,
		VehicleID:     vid,
		Processed:     0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := vehicleModel.InsertTestDrive(td)

	theData := JSONResponse{
		OK: true,
	}

	// build the json response from the struct
	out, err := json.MarshalIndent(theData, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// send json to client
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		errorLog.Println(err)
	}
}

// SendFriend sends to a friend
func SendFriend(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	interest := r.Form.Get("interested")
	url := r.Form.Get("url")

	content := fmt.Sprintf(`
		<p>
			Hi:
			<br>
			<br>
			%s thought you might be interested in this item at Jim Gilbert's PowerSports:
			<br><br>
			%s
			<br><br>
			You can see the item by following this link:
			<a href='http://%s'>Click here to see the item!</a>
		</p>
`, name, interest, url)

	mailMessage := maildata.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     fmt.Sprintf("%s thought you might be intersted in this item from Jim Gilbert's PowerSports", name),
		Content:     template.HTML(content),
		Template:    "generic-email.mail.tmpl",
	}

	helpers.SendEmail(mailMessage)

	theData := JSONResponse{
		OK: true,
	}

	// build the json response from the struct
	out, err := json.MarshalIndent(theData, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// send json to client
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		errorLog.Println(err)
	}
}

// CreditApp displays credit app page
func CreditApp(w http.ResponseWriter, r *http.Request) {
	pg, err := pageModel.GetBySlug("credit-application")

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rowSets := make(map[string]interface{})
	var years []int
	for y := time.Now().Year(); y > (time.Now().Year() - 100); y-- {
		years = append(years, y)
	}

	rowSets["years"] = years

	helpers.Render(w, r, "credit-app.page.tmpl", &templates.TemplateData{
		Form:    forms.New(nil),
		Page:    pg,
		RowSets: rowSets,
	})
}

// PostCreditApp handles ajax post of credit application
func PostCreditApp(w http.ResponseWriter, r *http.Request) {
	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "y", "m", "y", "phone", "address", "city", "province", "zip", "rent", "income", "vehicle", "g-recaptcha-response")

	form.RecaptchaValid(r.RemoteAddr)

	if !form.Valid() {
		theData := JSONResponse{
			OK:      false,
			Message: "Form error",
		}

		// build the json response from the struct
		out, err := json.MarshalIndent(theData, "", "    ")
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		// send json to client
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(out)
		if err != nil {
			errorLog.Println(err)
		}
		return
	}

	// create email
	content := fmt.Sprintf(`
		<p>
			<strong>PowerSports Credit Application</strong>:<br><br>
			<strong>Name:</strong> %s  %s<br>
			<strong>Date of birth:</strong> %s <br>
			<strong>Email:</strong> %s <br>
			<strong>Phone:</strong> %s <br>
			<strong>Address:</strong> %s %s, %s, %s<br>
			<strong>Rent/Mortgage</strong>: %s<br>
			<strong>Employer</strong>: %s<br>
			<strong>Income</strong>: %s<br>
			<strong>Interested In:</strong><br><br>
			%s
		</p>
`,
		form.Get("first_name"),
		form.Get("last_name"),
		fmt.Sprintf("%s-%s-%s", form.Get("y"), form.Get("m"), form.Get("d")),
		form.Get("phone"),
		form.Get("email"),
		form.Get("address"),
		form.Get("city"),
		form.Get("province"),
		form.Get("zip"),
		form.Get("rent"),
		form.Get("employer"),
		form.Get("income"),
		form.Get("vehicle"),
	)

	var cc []string
	cc = append(cc, "wheelsanddeals@pbssystems.com")
	cc = append(cc, "john.eliakis@wheelsanddeals.ca")
	cc = append(cc, "chelsea.gilbert@wheelsanddeals.ca")

	mailMessage := maildata.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     "PowerSports Credit Application",
		Content:     template.HTML(content),
		Template:    "generic-email.mail.tmpl",
		CC:          cc,
	}

	infoLog.Println("Sending email")

	helpers.SendEmail(mailMessage)

	// save the application
	creditApp := clientmodels.CreditApp{
		FirstName: form.Get("first_name"),
		LastName:  form.Get("last_name"),
		Email:     form.Get("email"),
		Phone:     form.Get("phone"),
		Address:   form.Get("address"),
		City:      form.Get("city"),
		Province:  form.Get("province"),
		Zip:       form.Get("zip"),
		Vehicle:   form.Get("vehicle"),
		Processed: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := vehicleModel.InsertCreditApp(creditApp)
	if err != nil {
		errorLog.Println(err)
	}

	theData := JSONResponse{
		OK: true,
	}

	// build the json response from the struct
	out, err := json.MarshalIndent(theData, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// send json to client
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		errorLog.Println(err)
	}

}
