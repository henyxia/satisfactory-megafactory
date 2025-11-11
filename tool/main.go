package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func downloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getItemList() {
	url := "https://satisfactory-calculator.com/en/items"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("HTTP error: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	id := 1
	doc.Find(".card-body").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("h6").Text())
		if name == "" {
			return
		}

		// Create Go variable name (remove spaces/special chars)
		varName := strings.ReplaceAll(name, " ", "")
		varName = strings.ReplaceAll(varName, "-", "")
		varName = strings.ReplaceAll(varName, "_", "")

		fmt.Printf("item%s := &models.Item{ID: %d, Name: \"%s\", Type: models.ItemTypeItem}\n", varName, id, name)
		fmt.Printf("items = append(items, item%s)\n", varName)
		id++
	})
}

func getItemAssets() {
	url := "https://satisfactory-calculator.com/en/items"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("HTTP error: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	doc.Find(".card-body").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("h6").Text())
		if name == "" {
			return
		}

		name = strings.ReplaceAll(name, " ", "-")
		name = strings.ToLower(name)
		name = fmt.Sprintf("item-%s.png", name)

		url, found := s.Find("img").Attr("src")
		if !found {
			return
		}

		fmt.Printf("%s = %s\n", name, url)
		downloadFile(url, name)
	})
}

/*
 * RECIPES
 */
type Ingredient struct {
	Name string  `json:"item"`
	Qt   float32 `json:"amount"`
}

func (i *Ingredient) String() string {
	return fmt.Sprintf("%0.1f * %s", i.Qt, i.Name)
}

type Recipe struct {
	ClassName   string       `json:"className"`
	Name        string       `json:"name"`
	Stable      bool         `json:"stable"`
	Duration    float32      `json:"duration"`
	ProducedIn  []string     `json:"producedIn"`
	Ingredients []Ingredient `json:"ingredients"`
	Products    []Ingredient `json:"products"`
	producedIn  string
}

/* source:
 * https://satisfactory.wiki.gg/wiki/Template:DocsRecipes.json?action=edit
 */
func extractRecipeList() {
	factoryConvertor := map[string]string{
		"Desc_AssemblerMk1_C":     "buildingAssember",
		"Desc_Blender_C":          "buildingBlender",
		"Desc_ConstructorMk1_C":   "buildingConstructor",
		"Desc_Converter_C":        "buildingConverter",
		"Desc_FoundryMk1_C":       "buildingFoundry",
		"Desc_GeneratorNuclear_C": "buildingNuclearGenerator",
		"Desc_HadronCollider_C":   "buildingParicleAccelerator",
		"Desc_ManufacturerMk1_C":  "buildingManufacurer",
		"Desc_OilRefinery_C":      "buildingRefinery",
		"Desc_Packager_C":         "buildingPackager",
		"Desc_QuantumEncoder_C":   "buildingQuantumEncoder",
		"Desc_SmelterMk1_C":       "buildingSmelter",
	}

	itemList := map[string]string{
		"Desc_Stone_C":                     "itemLimestone",
		"Desc_OreIron_C":                   "itemIronOre",
		"Desc_OreCopper_C":                 "itemCopperOre",
		"Desc_OreGold_C":                   "itemCateriumOre",
		"Desc_Coal_C":                      "itemCoal",
		"Desc_RawQuartz_C":                 "itemRawQuartz",
		"Desc_Sulfur_C":                    "itemSulfur",
		"Desc_OreBauxite_C":                "itemBauxite",
		"Desc_SAM_C":                       "itemSAM",
		"Desc_OreUranium_C":                "itemUranium",
		"Desc_IronIngot_C":                 "itemIronIngot",
		"Desc_CopperIngot_C":               "itemCopperIngot",
		"Desc_GoldIngot_C":                 "itemCateriumIngot",
		"Desc_SteelIngot_C":                "itemSteelIngot",
		"Desc_AluminumIngot_C":             "itemAluminumIngot",
		"Desc_FicsiteIngot_C":              "itemFicsiteIngot",
		"Desc_Cement_C":                    "itemConcrete",
		"Desc_QuartzCrystal_C":             "itemQuartzCrystal",
		"Desc_Silica_C":                    "itemSilica",
		"Desc_CopperDust_C":                "itemCopperPowder",
		"Desc_PolymerResin_C":              "itemPolymerResin",
		"Desc_PetroleumCoke_C":             "itemPetroleumCoke",
		"Desc_AluminumScrap_C":             "itemAluminumScrap",
		"Desc_AlienProtein_C":              "itemAlienProtein",
		"Desc_AlienDNACapsule_C":           "itemAlienDNACapsule",
		"Desc_Water_C":                     "itemWater",
		"Desc_LiquidOil_C":                 "itemCrudeOil",
		"Desc_HeavyOilResidue_C":           "itemHeavyOilResidue",
		"Desc_LiquidFuel_C":                "itemFuel",
		"Desc_LiquidBiofuel_C":             "itemLiquidBiofuel",
		"Desc_LiquidTurboFuel_C":           "itemTurbofuel",
		"Desc_AluminaSolution_C":           "itemAluminaSolution",
		"Desc_SulfuricAcid_C":              "itemSulfuricAcid",
		"Desc_NitricAcid_C":                "itemNitricAcid",
		"Desc_DissolvedSilica_C":           "itemDissolvedSilica",
		"Desc_NitrogenGas_C":               "itemNitrogenGas",
		"Desc_RocketFuel_C":                "itemRocketFuel",
		"Desc_IonizedFuel_C":               "itemIonizedFuel",
		"Desc_DarkEnergy_C":                "itemDarkMatterResidue",
		"Desc_QuantumEnergy_C":             "itemExcitedPhotonicMatter",
		"Desc_IronRod_C":                   "itemIronRod",
		"Desc_IronScrew_C":                 "itemScrews",
		"Desc_IronPlate_C":                 "itemIronPlate",
		"Desc_IronPlateReinforced_C":       "itemReinforcedIronPlate",
		"Desc_CopperSheet_C":               "itemCopperSheet",
		"Desc_AluminumPlate_C":             "itemAlcladAluminumSheet",
		"Desc_AluminumCasing_C":            "itemAluminumCasing",
		"Desc_SteelPipe_C":                 "itemSteelPipe",
		"Desc_SteelPlate_C":                "itemSteelBeam",
		"Desc_SteelPlateReinforced_C":      "itemEncasedIndustrialBeam",
		"Desc_ModularFrame_C":              "itemModularFrame",
		"Desc_ModularFrameHeavy_C":         "itemHeavyModularFrame",
		"Desc_ModularFrameFused_C":         "itemFusedModularFrame",
		"Desc_FicsiteMesh_C":               "itemFicsiteTrigon",
		"Desc_Fabric_C":                    "itemFabric",
		"Desc_Plastic_C":                   "itemPlastic",
		"Desc_Rubber_C":                    "itemRubber",
		"Desc_Rotor_C":                     "itemRotor",
		"Desc_Stator_C":                    "itemStator",
		"Desc_Battery_C":                   "itemBattery",
		"Desc_Motor_C":                     "itemMotor",
		"Desc_AluminumPlateReinforced_C":   "itemHeatSink",
		"Desc_CoolingSystem_C":             "itemCoolingSystem",
		"Desc_MotorLightweight_C":          "itemTurboMotor",
		"Desc_Wire_C":                      "itemWire",
		"Desc_Cable_C":                     "itemCable",
		"Desc_HighSpeedWire_C":             "itemQuickwire",
		"Desc_CircuitBoard_C":              "itemCircuitBoard",
		"Desc_CircuitBoardHighSpeed_C":     "itemAILimiter",
		"Desc_HighSpeedConnector_C":        "itemHighSpeedConnector",
		"Desc_SAMIngot_C":                  "itemReanimatedSAM",
		"Desc_SAMFluctuator_C":             "itemSAMFluctuator",
		"Desc_Computer_C":                  "itemComputer",
		"Desc_ComputerSuper_C":             "itemSupercomputer",
		"Desc_ModularFrameLightweight_C":   "itemRadioControlUnit",
		"Desc_CrystalOscillator_C":         "itemCrystalOscillator",
		"Desc_QuantumOscillator_C":         "itemSuperpositionOscillator",
		"Desc_Diamond_C":                   "itemDiamonds",
		"Desc_TimeCrystal_C":               "itemTimeCrystal",
		"Desc_DarkMatter_C":                "itemDarkMatterCrystal",
		"Desc_SingularityCell_C":           "itemSingularityCell",
		"Desc_TemporalProcessor_C":         "itemNeuralQuantumProcessor",
		"Desc_AlienPowerFuel_C":            "itemAlienPowerMatrix",
		"Desc_FluidCanister_C":             "itemEmptyCanister",
		"Desc_GasTank_C":                   "itemEmptyFluidTank",
		"Desc_PressureConversionCube_C":    "itemPressureConversionCube",
		"Desc_PackagedWater_C":             "itemPackagedWater",
		"Desc_PackagedAlumina_C":           "itemPackagedAluminaSolution",
		"Desc_PackagedSulfuricAcid_C":      "itemPackagedSulfuricAcid",
		"Desc_PackagedNitricAcid_C":        "itemPackagedNitricAcid",
		"Desc_PackagedNitrogenGas_C":       "itemPackagedNitrogenGas",
		"Desc_Leaves_C":                    "itemLeaves",
		"Desc_Mycelia_C":                   "itemMycelia",
		"Desc_FlowerPetals_C":              "itemFlowerPetals",
		"Desc_Wood_C":                      "itemWood",
		"Desc_GenericBiomass_C":            "itemBiomass",
		"Desc_CompactedCoal_C":             "itemCompactedCoal",
		"Desc_PackagedOil_C":               "itemPackagedOil",
		"Desc_PackagedOilResidue_C":        "itemPackagedHeavyOilResidue",
		"Desc_Biofuel_C":                   "itemSolidBiofuel",
		"Desc_Fuel_C":                      "itemPackagedFuel",
		"Desc_PackagedBiofuel_C":           "itemPackagedLiquidBiofuel",
		"Desc_TurboFuel_C":                 "itemPackagedTurbofuel",
		"Desc_PackagedRocketFuel_C":        "itemPackagedRocketFuel",
		"Desc_PackagedIonizedFuel_C":       "itemPackagedIonizedFuel",
		"Desc_NuclearFuelRod_C":            "itemUraniumFuelRod",
		"Desc_PlutoniumFuelRod_C":          "itemPlutoniumFuelRod",
		"Desc_Gunpowder_C":                 "itemBlackPowder",
		"Desc_GunpowderMK2_C":              "itemSmokelessPowder",
		"Desc_Filter_C":                    "itemGasFilter",
		"Desc_ColorCartridge_C":            "itemColorCartridge",
		"Desc_Beacon_C":                    "itemBeacon",
		"Desc_HazmatFilter_C":              "itemIodineInfusedFilter",
		"Desc_SpikedRebar_C":               "itemIronRebar",
		"Desc_Rebar_Stunshot_C":            "itemStunRebar",
		"Desc_Rebar_Spreadshot_C":          "itemShatterRebar",
		"Desc_Rebar_Explosive_C":           "itemExplosiveRebar",
		"Desc_CartridgeStandard_C":         "itemRifleAmmo",
		"Desc_CartridgeSmartProjectile_C":  "itemHomingRifleAmmo",
		"Desc_CartridgeChaos_C":            "itemTurboRifleAmmo",
		"Desc_NobeliskExplosive_C":         "itemNobelisk",
		"Desc_NobeliskGas_C":               "itemGasNobelisk",
		"Desc_NobeliskShockwave_C":         "itemPulseNobelisk",
		"Desc_NobeliskCluster_C":           "itemClusterNobelisk",
		"Desc_NobeliskNuke_C":              "itemNukeNobelisk",
		"Desc_ElectromagneticControlRod_C": "itemElectromagneticControlRod",
		"Desc_UraniumCell_C":               "itemEncasedUraniumCell",
		"Desc_NonFissibleUranium_C":        "itemNonFissileUranium",
		"Desc_PlutoniumPellet_C":           "itemPlutoniumPellet",
		"Desc_PlutoniumCell_C":             "itemEncasedPlutoniumCell",
		"Desc_Ficsonium_C":                 "itemFicsonium",
		"Desc_FicsoniumFuelRod_C":          "itemFicsoniumFuelRod",
		"Desc_NuclearWaste_C":              "itemUraniumWaste",
		"Desc_PlutoniumWaste_C":            "itemPlutoniumWaste",
		"Desc_Crystal_C":                   "itemBluePowerSlug",
		"Desc_Crystal_mk2_C":               "itemYellowPowerSlug",
		"Desc_Crystal_mk3_C":               "itemPurplePowerSlug",
		"Desc_CrystalShard_C":              "itemPowerShard",
		"Desc_FICSITCoupon_C":              "itemFICSITCoupon",
		"Desc_SpaceElevatorPart_1_C":       "itemSmartPlating",
		"Desc_SpaceElevatorPart_2_C":       "itemVersatileFramework",
		"Desc_SpaceElevatorPart_3_C":       "itemAutomatedWiring",
		"Desc_SpaceElevatorPart_4_C":       "itemModularEngine",
		"Desc_SpaceElevatorPart_5_C":       "itemAdaptiveControlUnit",
		"Desc_SpaceElevatorPart_6_C":       "itemAssemblyDirectorSystem",
		"Desc_SpaceElevatorPart_7_C":       "itemMagneticFieldGenerator",
		"Desc_SpaceElevatorPart_8_C":       "itemThermalPropulsionRocket",
		"Desc_SpaceElevatorPart_9_C":       "itemNuclearPasta",
		"Desc_SpaceElevatorPart_10_C":      "itemBiochemicalSculptor",
		"Desc_SpaceElevatorPart_11_C":      "itemBallisticWarpDrive",
		"Desc_SpaceElevatorPart_12_C":      "itemAIExpansionServer",
		"BP_ItemDescriptorPortableMiner_C": "itemPortableMiner",
		"Desc_HatcherParts_C":              "itemHatcherProtein",
		"Desc_HogParts_C":                  "itemHogProtein",
		"Desc_SpitterParts_C":              "itemSpitterProtein",
		"Desc_StingerParts_C":              "itemStingerProtein",
	}

	ficsmas := map[string]bool{
		"Blue FICSMAS Ornament":   true,
		"Copper FICSMAS Ornament": true,
		"FICSMAS Actual Snow":     true,
		"FICSMAS Bow":             true,
		"FICSMAS Ornament Bundle": true,
		"FICSMAS Tree Branch":     true,
		"FICSMAS Wonder Star":     true,
		"FICSMAS Wreath":          true,
		"Iron FICSMAS Ornament":   true,
		"Red FICSMAS Ornament":    true,
		"Candy Cane":              true,
		"Sparkly Fireworks":       true,
		"Snowball":                true,
		"Sweet Fireworks":         true,
		"Fancy Fireworks":         true,
	}

	recipes := make(map[string][]Recipe)

	// read file
	recipeFile, err := os.Open("DocsRecipes.json")
	if err != nil {
		fmt.Println("opening config file", err.Error())
		return
	}

	jsonParser := json.NewDecoder(recipeFile)
	if err = jsonParser.Decode(&recipes); err != nil {
		fmt.Println("parsing config file", err.Error())
		return
	}

	fmt.Println("package main")
	fmt.Println("")
	fmt.Println("import \"github.com/henyxia/satisfactory-megafactory/models\"")
	fmt.Println("")
	fmt.Println("func (cfg *config) setRecipe() {")
	var rendered []string

	for _, _recipe := range recipes {
		for _, recipe := range _recipe {
			if len(recipe.ProducedIn) == 0 {
				continue
			}
			if !recipe.Stable {
				continue
			}
			_, isFicsmas := ficsmas[recipe.Name]
			if isFicsmas {
				continue
			}
			_, ok := factoryConvertor[recipe.ProducedIn[0]]
			if ok {
				recipe.producedIn = factoryConvertor[recipe.ProducedIn[0]]
			} else {
				fmt.Println("key " + recipe.ProducedIn[0] + "not found!")
				continue
			}

			var input []string
			for _, ingredient := range recipe.Ingredients {
				input = append(input, fmt.Sprintf(
					"{Item: *%s, Qt: %d}",
					itemList[ingredient.Name],
					uint(ingredient.Qt*10),
				))
			}

			var output []string
			for _, product := range recipe.Products {
				output = append(output, fmt.Sprintf(
					"{Item: *%s, Qt: %d}",
					itemList[product.Name],
					uint(product.Qt*10),
				))
			}

			slug := strings.ToLower(recipe.Name)
			slug = strings.ReplaceAll(slug, " ", "_")
			slug = strings.ReplaceAll(slug, "(", "")
			slug = strings.ReplaceAll(slug, ")", "")

			rendered = append(rendered, fmt.Sprintf(
				"		{Slug:\"%s\", Name:\"%s\", Duration: %d, ProducedIn: %s, Input: []models.RecipeInput{%s}, Output: []models.RecipeOutput{%s}},",
				slug,
				recipe.Name,
				uint(recipe.Duration*10),
				recipe.producedIn,
				strings.Join(input, ", "),
				strings.Join(output, ", "),
			))
		}
	}

	slices.Sort(rendered)
	fmt.Println("	recipes := []models.Recipe{")
	for _, line := range rendered {
		fmt.Println(line)
	}
	fmt.Println("	}")
}

/*
 * MAIN
 */

func usage(error string) {
	if error != "" {
		fmt.Println("error: " + error)
		fmt.Println("")
	}
	fmt.Println("usage: ./tool command")
	fmt.Println("")
	fmt.Println("	get-item-list			return the list of items")
	fmt.Println("	get-item-assets			download all item assets")
	fmt.Println("	extract-recipe-list		read the DocsRecipes.json file and return the recipe list")
}

func main() {
	if len(os.Args) != 2 {
		usage("missing command")
		return
	}

	switch os.Args[1] {
	case "get-item-list":
		getItemList()
	case "get-item-assets":
		getItemAssets()
	case "extract-recipe-list":
		extractRecipeList()
	default:
		usage("no command specified")
	}
}
