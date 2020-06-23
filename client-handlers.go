package clienthandlers

import (
	"encoding/json"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	channel_data "github.com/tsawler/goblender/pkg/channel-data"
	"github.com/tsawler/goblender/pkg/forms"
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/templates"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var vehicleModel *clientdb.VehicleModel

const (
	SOLD    = 0
	FORSALE = 1
	PENDING = 2
	TRADEIN = 3
)

const (
	All           = 0
	ATVBruteForce = 8
	ATVMule       = 11
	ATVTeryx      = 12
	Car           = 1
	ElectricBike  = 16
	JetSki        = 13
	Mercury       = 10
	Motorcycle    = 7
	Other         = 3
	PontoonBoat   = 9
	PowerBoat     = 15
	Scooter       = 17
	SUV           = 5
	Trailer       = 14
	Truck         = 2
	MiniVan       = 6
	Unknown       = 4
)

// JSONResponse is a generic struct to hold json responses
type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// CompareVehicles Show 2 or 3 vehicles in table TODO
func CompareVehicles(w http.ResponseWriter, r *http.Request) {
	idString := r.Form.Get("ids")
	infoLog.Println("Ids:", idString)

	ids := strings.Split(idString, ",")
	var items []clientmodels.Vehicle

	for _, x := range ids {
		infoLog.Println("ID:", x)
		vid, _ := strconv.Atoi(x)
		v, _ := vehicleModel.GetPowerSportItem(vid)
		items = append(items, v)
	}

	rowSets := make(map[string]interface{})
	rowSets["items"] = items
	helpers.Render(w, r, "compare.page.tmpl", &templates.TemplateData{
		RowSets: rowSets,
	})
}

// GetAllMotorcycles gets all motorcycles
func GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/motorcyclesforsale/fredericton"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 7
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "motorcycle-inventory")
}

// GetAllBruteForce gets all brute force
func GetAllBruteForce(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/atvs/brute-force-inventory"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 8
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-brute-force")
}

// GetAllTeryx gets all teryx
func GetAllTeryx(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/sidexsides/teryx-for-sale"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 12
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-teryx")
}

// GetAllMule gets all mules
func GetAllMule(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/utvs/kawasaki-mules-for-sale"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 11
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-mule")
}

// GetAllJetSki gets all new jetskis
func GetAllJetSki(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/personalwatercraft/new-jetskis-forsale"
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
	stringMap["pager-url"] = "/ebikes/pedego-inventory"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 16
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "electric-bikes")
}

// GetAllScooters gets all scooters
func GetAllScooters(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/scooters-mopeds/vespa-piaggio-gas-electric"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 17
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "vespa-piaggio-gas-and-electric-scooters-mopeds")
}

// GetAllPontoonBoats gets all pontoon boats (bennington & crestliner
func GetAllPontoonBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/pontoon-boats/bennington-crestliner-dealer"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 9
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "pontoon-boats-bennington-and-crestliner")
}

// GetAllPowerBoats gets all power boats
func GetAllPowerBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/boatsforsale/fishing-speed-jon-aluminum-power-glastron"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 15
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "speed-boats-aluminum-boats-power-boats")
}

// GetAllUsedPowerSports gets all used powersports
func GetAllUsedPowerSports(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/used/used-motorbikes-atvs-wheelers-trailers"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 1000
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "used-motorcycles-used-atv-used-boats-used-pontoons")
}

// GetAllUsedBoats gets all used boats
func GetAllUsedBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/used-boats-for-sale/pontoons-pwc-jetski-fishboats-boattrailers"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 1001
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "used-boats")
}

// GetAllTrailers gets all used powersports
func GetAllTrailers(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/newtrailers/boat-stv-utility"
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

	pg, err := repo.DB.GetPageBySlug(slug)
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

	pg, err := repo.DB.GetPageBySlug("power-sports-item")
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

	if item.Status != FORSALE {
		// item is sold, or whatever
		http.Redirect(w, r, "/motorcyclesforsale/fredericton", http.StatusMovedPermanently)
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
	vid, _ := strconv.Atoi(r.Form.Get("vehicle_id"))
	msg := r.Form.Get("msg")
	if msg != "" {
		interest = msg
	}

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
	cc = append(cc, "wheelsanddealspowersports@pbssystems.com")

	mailMessage := channel_data.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     "PowerSports Quick Quote Request",
		Content:     template.HTML(content),
		Template:    "bootstrap.mail.tmpl",
		CC:          cc,
	}

	helpers.SendEmail(mailMessage)

	qq := clientmodels.QuickQuote{
		UsersName: name,
		Email:     email,
		Phone:     phone,
		VehicleID: vid,
		Processed: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := vehicleModel.InsertQuickQuote(qq)
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
	cc = append(cc, "wheelsanddealspowersports@pbssystems.com")
	//cc = append(cc, "john.eliakis@wheelsanddeals.ca")

	mailMessage := channel_data.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     "PowerSports Test Drive Request",
		Content:     template.HTML(content),
		Template:    "bootstrap.mail.tmpl",
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

// SendFriend sends to a friend
func SendFriend(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	interest := r.Form.Get("interested")
	url := r.Form.Get("url")
	infoLog.Println("url is ", url)

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
			<a href='%s'>Click here to see the item!</a>
		</p>
`, name, interest, url)

	mailMessage := channel_data.MailData{
		ToName:      "",
		ToAddress:   email,
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     fmt.Sprintf("%s thought you might be intersted in this item from Jim Gilbert's PowerSports", name),
		Content:     template.HTML(content),
		Template:    "bootstrap.mail.tmpl",
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
	pg, err := repo.DB.GetPageBySlug("credit-application")

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
	cc = append(cc, "wheelsanddealspowersports@pbssystems.com")
	//cc = append(cc, "john.eliakis@wheelsanddeals.ca")
	cc = append(cc, "chelsea.gilbert@wheelsanddeals.ca")

	mailMessage := channel_data.MailData{
		ToName:      "",
		ToAddress:   "alex.gilbert@wheelsanddeals.ca",
		FromName:    app.PreferenceMap["smtp-from-name"],
		FromAddress: app.PreferenceMap["smtp-from-email"],
		Subject:     "PowerSports Credit Application",
		Content:     template.HTML(content),
		Template:    "bootstrap.mail.tmpl",
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
