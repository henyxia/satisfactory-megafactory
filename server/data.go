package server

import (
	"github.com/henyxia/satisfactory-megafactory/models"
)

func (s *Server) setData() error {
	var buildings []models.Building
	buildingAssember := models.Building{ID: 1, Name: "Assember"}
	buildingBlender := models.Building{ID: 2, Name: "Blender"}
	buildingConstructor := models.Building{ID: 3, Name: "Constructor"}
	buildingConverter := models.Building{ID: 4, Name: "Converter"}
	buildingFoundry := models.Building{ID: 5, Name: "Foundry"}
	buildingNuclearGenerator := models.Building{ID: 6, Name: "Nuclear Generator"}
	buildingParicleAccelerator := models.Building{ID: 7, Name: "Particle Accelerator"}
	buildingManufacurer := models.Building{ID: 8, Name: "Manufacturer"}
	buildingRefinery := models.Building{ID: 9, Name: "Refinery"}
	buildingPackager := models.Building{ID: 10, Name: "Packager"}
	buildingQuantumEncoder := models.Building{ID: 11, Name: "Quantum Encoder"}
	buildingSmelter := models.Building{ID: 12, Name: "Smelter"}
	buildings = append(buildings,
		buildingAssember,
		buildingBlender,
		buildingConstructor,
		buildingConverter,
		buildingFoundry,
		buildingNuclearGenerator,
		buildingParicleAccelerator,
		buildingManufacurer,
		buildingRefinery,
		buildingPackager,
		buildingQuantumEncoder,
		buildingSmelter,
	)

	err := s.DB.Save(buildings).Error
	if err != nil {
		return err
	}

	var items []*models.Item
	itemLimestone := &models.Item{ID: 1, Name: "Limestone", Type: models.ItemTypeItem}
	itemIronOre := &models.Item{ID: 2, Name: "Iron Ore", Type: models.ItemTypeItem}
	itemCopperOre := &models.Item{ID: 3, Name: "Copper Ore", Type: models.ItemTypeItem}
	itemCateriumOre := &models.Item{ID: 4, Name: "Caterium Ore", Type: models.ItemTypeItem}
	itemCoal := &models.Item{ID: 5, Name: "Coal", Type: models.ItemTypeItem}
	itemRawQuartz := &models.Item{ID: 6, Name: "Raw Quartz", Type: models.ItemTypeItem}
	itemSulfur := &models.Item{ID: 7, Name: "Sulfur", Type: models.ItemTypeItem}
	itemBauxite := &models.Item{ID: 8, Name: "Bauxite", Type: models.ItemTypeItem}
	itemSAM := &models.Item{ID: 9, Name: "SAM", Type: models.ItemTypeItem}
	itemUranium := &models.Item{ID: 10, Name: "Uranium", Type: models.ItemTypeItem}
	itemIronIngot := &models.Item{ID: 11, Name: "Iron Ingot", Type: models.ItemTypeItem}
	itemCopperIngot := &models.Item{ID: 12, Name: "Copper Ingot", Type: models.ItemTypeItem}
	itemCateriumIngot := &models.Item{ID: 13, Name: "Caterium Ingot", Type: models.ItemTypeItem}
	itemSteelIngot := &models.Item{ID: 14, Name: "Steel Ingot", Type: models.ItemTypeItem}
	itemAluminumIngot := &models.Item{ID: 15, Name: "Aluminum Ingot", Type: models.ItemTypeItem}
	itemFicsiteIngot := &models.Item{ID: 16, Name: "Ficsite Ingot", Type: models.ItemTypeItem}
	itemConcrete := &models.Item{ID: 17, Name: "Concrete", Type: models.ItemTypeItem}
	itemQuartzCrystal := &models.Item{ID: 18, Name: "Quartz Crystal", Type: models.ItemTypeItem}
	itemSilica := &models.Item{ID: 19, Name: "Silica", Type: models.ItemTypeItem}
	itemCopperPowder := &models.Item{ID: 20, Name: "Copper Powder", Type: models.ItemTypeItem}
	itemPolymerResin := &models.Item{ID: 21, Name: "Polymer Resin", Type: models.ItemTypeItem}
	itemPetroleumCoke := &models.Item{ID: 22, Name: "Petroleum Coke", Type: models.ItemTypeItem}
	itemAluminumScrap := &models.Item{ID: 23, Name: "Aluminum Scrap", Type: models.ItemTypeItem}
	itemAlienProtein := &models.Item{ID: 24, Name: "Alien Protein", Type: models.ItemTypeItem}
	itemAlienDNACapsule := &models.Item{ID: 25, Name: "Alien DNA Capsule", Type: models.ItemTypeItem}
	itemWater := &models.Item{ID: 26, Name: "Water", Type: models.ItemTypeLiquid}
	itemCrudeOil := &models.Item{ID: 27, Name: "Crude Oil", Type: models.ItemTypeLiquid}
	itemHeavyOilResidue := &models.Item{ID: 28, Name: "Heavy Oil Residue", Type: models.ItemTypeLiquid}
	itemFuel := &models.Item{ID: 29, Name: "Fuel", Type: models.ItemTypeLiquid}
	itemLiquidBiofuel := &models.Item{ID: 30, Name: "Liquid Biofuel", Type: models.ItemTypeLiquid}
	itemTurbofuel := &models.Item{ID: 31, Name: "Turbofuel", Type: models.ItemTypeLiquid}
	itemAluminaSolution := &models.Item{ID: 32, Name: "Alumina Solution", Type: models.ItemTypeLiquid}
	itemSulfuricAcid := &models.Item{ID: 33, Name: "Sulfuric Acid", Type: models.ItemTypeLiquid}
	itemNitricAcid := &models.Item{ID: 34, Name: "Nitric Acid", Type: models.ItemTypeLiquid}
	itemDissolvedSilica := &models.Item{ID: 35, Name: "Dissolved Silica", Type: models.ItemTypeLiquid}
	itemNitrogenGas := &models.Item{ID: 36, Name: "Nitrogen Gas", Type: models.ItemTypeGas}
	itemRocketFuel := &models.Item{ID: 37, Name: "Rocket Fuel", Type: models.ItemTypeGas}
	itemIonizedFuel := &models.Item{ID: 38, Name: "Ionized Fuel", Type: models.ItemTypeGas}
	itemDarkMatterResidue := &models.Item{ID: 39, Name: "Dark Matter Residue", Type: models.ItemTypeGas}
	itemExcitedPhotonicMatter := &models.Item{ID: 40, Name: "Excited Photonic Matter", Type: models.ItemTypeGas}
	itemIronRod := &models.Item{ID: 41, Name: "Iron Rod", Type: models.ItemTypeItem}
	itemScrews := &models.Item{ID: 42, Name: "Screws", Type: models.ItemTypeItem}
	itemIronPlate := &models.Item{ID: 43, Name: "Iron Plate", Type: models.ItemTypeItem}
	itemReinforcedIronPlate := &models.Item{ID: 44, Name: "Reinforced Iron Plate", Type: models.ItemTypeItem}
	itemCopperSheet := &models.Item{ID: 45, Name: "Copper Sheet", Type: models.ItemTypeItem}
	itemAlcladAluminumSheet := &models.Item{ID: 46, Name: "Alclad Aluminum Sheet", Type: models.ItemTypeItem}
	itemAluminumCasing := &models.Item{ID: 47, Name: "Aluminum Casing", Type: models.ItemTypeItem}
	itemSteelPipe := &models.Item{ID: 48, Name: "Steel Pipe", Type: models.ItemTypeItem}
	itemSteelBeam := &models.Item{ID: 49, Name: "Steel Beam", Type: models.ItemTypeItem}
	itemEncasedIndustrialBeam := &models.Item{ID: 50, Name: "Encased Industrial Beam", Type: models.ItemTypeItem}
	itemModularFrame := &models.Item{ID: 51, Name: "Modular Frame", Type: models.ItemTypeItem}
	itemHeavyModularFrame := &models.Item{ID: 52, Name: "Heavy Modular Frame", Type: models.ItemTypeItem}
	itemFusedModularFrame := &models.Item{ID: 53, Name: "Fused Modular Frame", Type: models.ItemTypeItem}
	itemFicsiteTrigon := &models.Item{ID: 54, Name: "Ficsite Trigon", Type: models.ItemTypeItem}
	itemFabric := &models.Item{ID: 55, Name: "Fabric", Type: models.ItemTypeItem}
	itemPlastic := &models.Item{ID: 56, Name: "Plastic", Type: models.ItemTypeItem}
	itemRubber := &models.Item{ID: 57, Name: "Rubber", Type: models.ItemTypeItem}
	itemRotor := &models.Item{ID: 58, Name: "Rotor", Type: models.ItemTypeItem}
	itemStator := &models.Item{ID: 59, Name: "Stator", Type: models.ItemTypeItem}
	itemBattery := &models.Item{ID: 60, Name: "Battery", Type: models.ItemTypeItem}
	itemMotor := &models.Item{ID: 61, Name: "Motor", Type: models.ItemTypeItem}
	itemHeatSink := &models.Item{ID: 62, Name: "Heat Sink", Type: models.ItemTypeItem}
	itemCoolingSystem := &models.Item{ID: 63, Name: "Cooling System", Type: models.ItemTypeItem}
	itemTurboMotor := &models.Item{ID: 64, Name: "Turbo Motor", Type: models.ItemTypeItem}
	itemWire := &models.Item{ID: 65, Name: "Wire", Type: models.ItemTypeItem}
	itemCable := &models.Item{ID: 66, Name: "Cable", Type: models.ItemTypeItem}
	itemQuickwire := &models.Item{ID: 67, Name: "Quickwire", Type: models.ItemTypeItem}
	itemCircuitBoard := &models.Item{ID: 68, Name: "Circuit Board", Type: models.ItemTypeItem}
	itemAILimiter := &models.Item{ID: 69, Name: "AI Limiter", Type: models.ItemTypeItem}
	itemHighSpeedConnector := &models.Item{ID: 70, Name: "High-Speed Connector", Type: models.ItemTypeItem}
	itemReanimatedSAM := &models.Item{ID: 71, Name: "Reanimated SAM", Type: models.ItemTypeItem}
	itemSAMFluctuator := &models.Item{ID: 72, Name: "SAM Fluctuator", Type: models.ItemTypeItem}
	itemComputer := &models.Item{ID: 73, Name: "Computer", Type: models.ItemTypeItem}
	itemSupercomputer := &models.Item{ID: 74, Name: "Supercomputer", Type: models.ItemTypeItem}
	itemRadioControlUnit := &models.Item{ID: 75, Name: "Radio Control Unit", Type: models.ItemTypeItem}
	itemCrystalOscillator := &models.Item{ID: 76, Name: "Crystal Oscillator", Type: models.ItemTypeItem}
	itemSuperpositionOscillator := &models.Item{ID: 77, Name: "Superposition Oscillator", Type: models.ItemTypeItem}
	itemDiamonds := &models.Item{ID: 78, Name: "Diamonds", Type: models.ItemTypeItem}
	itemTimeCrystal := &models.Item{ID: 79, Name: "Time Crystal", Type: models.ItemTypeItem}
	itemDarkMatterCrystal := &models.Item{ID: 80, Name: "Dark Matter Crystal", Type: models.ItemTypeItem}
	itemSingularityCell := &models.Item{ID: 81, Name: "Singularity Cell", Type: models.ItemTypeItem}
	itemNeuralQuantumProcessor := &models.Item{ID: 82, Name: "Neural-Quantum Processor", Type: models.ItemTypeItem}
	itemAlienPowerMatrix := &models.Item{ID: 83, Name: "Alien Power Matrix", Type: models.ItemTypeItem}
	itemEmptyCanister := &models.Item{ID: 84, Name: "Empty Canister", Type: models.ItemTypeItem}
	itemEmptyFluidTank := &models.Item{ID: 85, Name: "Empty Fluid Tank", Type: models.ItemTypeItem}
	itemPressureConversionCube := &models.Item{ID: 86, Name: "Pressure Conversion Cube", Type: models.ItemTypeItem}
	itemPackagedWater := &models.Item{ID: 87, Name: "Packaged Water", Type: models.ItemTypeItem}
	itemPackagedAluminaSolution := &models.Item{ID: 88, Name: "Packaged Alumina Solution", Type: models.ItemTypeItem}
	itemPackagedSulfuricAcid := &models.Item{ID: 89, Name: "Packaged Sulfuric Acid", Type: models.ItemTypeItem}
	itemPackagedNitricAcid := &models.Item{ID: 90, Name: "Packaged Nitric Acid", Type: models.ItemTypeItem}
	itemPackagedNitrogenGas := &models.Item{ID: 91, Name: "Packaged Nitrogen Gas", Type: models.ItemTypeItem}
	itemLeaves := &models.Item{ID: 92, Name: "Leaves", Type: models.ItemTypeItem}
	itemMycelia := &models.Item{ID: 93, Name: "Mycelia", Type: models.ItemTypeItem}
	itemFlowerPetals := &models.Item{ID: 94, Name: "Flower Petals", Type: models.ItemTypeItem}
	itemWood := &models.Item{ID: 95, Name: "Wood", Type: models.ItemTypeItem}
	itemBiomass := &models.Item{ID: 96, Name: "Biomass", Type: models.ItemTypeItem}
	itemCompactedCoal := &models.Item{ID: 97, Name: "Compacted Coal", Type: models.ItemTypeItem}
	itemPackagedOil := &models.Item{ID: 98, Name: "Packaged Oil", Type: models.ItemTypeItem}
	itemPackagedHeavyOilResidue := &models.Item{ID: 99, Name: "Packaged Heavy Oil Residue", Type: models.ItemTypeItem}
	itemSolidBiofuel := &models.Item{ID: 100, Name: "Solid Biofuel", Type: models.ItemTypeItem}
	itemPackagedFuel := &models.Item{ID: 101, Name: "Packaged Fuel", Type: models.ItemTypeItem}
	itemPackagedLiquidBiofuel := &models.Item{ID: 102, Name: "Packaged Liquid Biofuel", Type: models.ItemTypeItem}
	itemPackagedTurbofuel := &models.Item{ID: 103, Name: "Packaged Turbofuel", Type: models.ItemTypeItem}
	itemPackagedRocketFuel := &models.Item{ID: 104, Name: "Packaged Rocket Fuel", Type: models.ItemTypeItem}
	itemPackagedIonizedFuel := &models.Item{ID: 105, Name: "Packaged Ionized Fuel", Type: models.ItemTypeItem}
	itemUraniumFuelRod := &models.Item{ID: 106, Name: "Uranium Fuel Rod", Type: models.ItemTypeItem}
	itemPlutoniumFuelRod := &models.Item{ID: 107, Name: "Plutonium Fuel Rod", Type: models.ItemTypeItem}
	itemBlackPowder := &models.Item{ID: 108, Name: "Black Powder", Type: models.ItemTypeItem}
	itemSmokelessPowder := &models.Item{ID: 109, Name: "Smokeless Powder", Type: models.ItemTypeItem}
	itemGasFilter := &models.Item{ID: 110, Name: "Gas Filter", Type: models.ItemTypeItem}
	itemColorCartridge := &models.Item{ID: 111, Name: "Color Cartridge", Type: models.ItemTypeItem}
	itemBeacon := &models.Item{ID: 112, Name: "Beacon", Type: models.ItemTypeItem}
	itemIodineInfusedFilter := &models.Item{ID: 113, Name: "Iodine-Infused Filter", Type: models.ItemTypeItem}
	itemIronRebar := &models.Item{ID: 114, Name: "Iron Rebar", Type: models.ItemTypeItem}
	itemStunRebar := &models.Item{ID: 115, Name: "Stun Rebar", Type: models.ItemTypeItem}
	itemShatterRebar := &models.Item{ID: 116, Name: "Shatter Rebar", Type: models.ItemTypeItem}
	itemExplosiveRebar := &models.Item{ID: 117, Name: "Explosive Rebar", Type: models.ItemTypeItem}
	itemRifleAmmo := &models.Item{ID: 118, Name: "Rifle Ammo", Type: models.ItemTypeItem}
	itemHomingRifleAmmo := &models.Item{ID: 119, Name: "Homing Rifle Ammo", Type: models.ItemTypeItem}
	itemTurboRifleAmmo := &models.Item{ID: 120, Name: "Turbo Rifle Ammo", Type: models.ItemTypeItem}
	itemNobelisk := &models.Item{ID: 121, Name: "Nobelisk", Type: models.ItemTypeItem}
	itemGasNobelisk := &models.Item{ID: 122, Name: "Gas Nobelisk", Type: models.ItemTypeItem}
	itemPulseNobelisk := &models.Item{ID: 123, Name: "Pulse Nobelisk", Type: models.ItemTypeItem}
	itemClusterNobelisk := &models.Item{ID: 124, Name: "Cluster Nobelisk", Type: models.ItemTypeItem}
	itemNukeNobelisk := &models.Item{ID: 125, Name: "Nuke Nobelisk", Type: models.ItemTypeItem}
	itemElectromagneticControlRod := &models.Item{ID: 126, Name: "Electromagnetic Control Rod", Type: models.ItemTypeItem}
	itemEncasedUraniumCell := &models.Item{ID: 127, Name: "Encased Uranium Cell", Type: models.ItemTypeItem}
	itemNonFissileUranium := &models.Item{ID: 128, Name: "Non-Fissile Uranium", Type: models.ItemTypeItem}
	itemPlutoniumPellet := &models.Item{ID: 129, Name: "Plutonium Pellet", Type: models.ItemTypeItem}
	itemEncasedPlutoniumCell := &models.Item{ID: 130, Name: "Encased Plutonium Cell", Type: models.ItemTypeItem}
	itemFicsonium := &models.Item{ID: 131, Name: "Ficsonium", Type: models.ItemTypeItem}
	itemFicsoniumFuelRod := &models.Item{ID: 132, Name: "Ficsonium Fuel Rod", Type: models.ItemTypeItem}
	itemUraniumWaste := &models.Item{ID: 133, Name: "Uranium Waste", Type: models.ItemTypeItem}
	itemPlutoniumWaste := &models.Item{ID: 134, Name: "Plutonium Waste", Type: models.ItemTypeItem}
	itemBluePowerSlug := &models.Item{ID: 135, Name: "Blue Power Slug", Type: models.ItemTypeItem}
	itemYellowPowerSlug := &models.Item{ID: 136, Name: "Yellow Power Slug", Type: models.ItemTypeItem}
	itemPurplePowerSlug := &models.Item{ID: 137, Name: "Purple Power Slug", Type: models.ItemTypeItem}
	itemPowerShard := &models.Item{ID: 138, Name: "Power Shard", Type: models.ItemTypeItem}
	itemFICSITCoupon := &models.Item{ID: 139, Name: "FICSIT Coupon", Type: models.ItemTypeItem}
	itemSmartPlating := &models.Item{ID: 140, Name: "Smart Plating", Type: models.ItemTypeItem}
	itemVersatileFramework := &models.Item{ID: 141, Name: "Versatile Framework", Type: models.ItemTypeItem}
	itemAutomatedWiring := &models.Item{ID: 142, Name: "Automated Wiring", Type: models.ItemTypeItem}
	itemModularEngine := &models.Item{ID: 143, Name: "Modular Engine", Type: models.ItemTypeItem}
	itemAdaptiveControlUnit := &models.Item{ID: 144, Name: "Adaptive Control Unit", Type: models.ItemTypeItem}
	itemAssemblyDirectorSystem := &models.Item{ID: 145, Name: "Assembly Director System", Type: models.ItemTypeItem}
	itemMagneticFieldGenerator := &models.Item{ID: 146, Name: "Magnetic Field Generator", Type: models.ItemTypeItem}
	itemThermalPropulsionRocket := &models.Item{ID: 147, Name: "Thermal Propulsion Rocket", Type: models.ItemTypeItem}
	itemNuclearPasta := &models.Item{ID: 148, Name: "Nuclear Pasta", Type: models.ItemTypeItem}
	itemBiochemicalSculptor := &models.Item{ID: 149, Name: "Biochemical Sculptor", Type: models.ItemTypeItem}
	itemBallisticWarpDrive := &models.Item{ID: 150, Name: "Ballistic Warp Drive", Type: models.ItemTypeItem}
	itemAIExpansionServer := &models.Item{ID: 151, Name: "AI Expansion Server", Type: models.ItemTypeItem}
	itemPortableMiner := &models.Item{ID: 152, Name: "Portable Miner", Type: models.ItemTypeItem}
	itemHatcherProtein := &models.Item{ID: 153, Name: "Hatcher Protein", Type: models.ItemTypeItem}
	itemHogProtein := &models.Item{ID: 154, Name: "Hog Protein", Type: models.ItemTypeItem}
	itemSpitterProtein := &models.Item{ID: 155, Name: "Spitter Protein", Type: models.ItemTypeItem}
	itemStingerProtein := &models.Item{ID: 156, Name: "Stinger Protein", Type: models.ItemTypeItem}

	// add items in list
	items = append(items,
		itemLimestone,
		itemIronOre,
		itemCopperOre,
		itemCateriumOre,
		itemCoal,
		itemRawQuartz,
		itemSulfur,
		itemBauxite,
		itemSAM,
		itemUranium,
		itemIronIngot,
		itemCopperIngot,
		itemCateriumIngot,
		itemSteelIngot,
		itemAluminumIngot,
		itemFicsiteIngot,
		itemConcrete,
		itemQuartzCrystal,
		itemSilica,
		itemCopperPowder,
		itemPolymerResin,
		itemPetroleumCoke,
		itemAluminumScrap,
		itemAlienProtein,
		itemAlienDNACapsule,
		itemWater,
		itemCrudeOil,
		itemHeavyOilResidue,
		itemFuel,
		itemLiquidBiofuel,
		itemTurbofuel,
		itemAluminaSolution,
		itemSulfuricAcid,
		itemNitricAcid,
		itemDissolvedSilica,
		itemNitrogenGas,
		itemRocketFuel,
		itemIonizedFuel,
		itemDarkMatterResidue,
		itemExcitedPhotonicMatter,
		itemIronRod,
		itemScrews,
		itemIronPlate,
		itemReinforcedIronPlate,
		itemCopperSheet,
		itemAlcladAluminumSheet,
		itemAluminumCasing,
		itemSteelPipe,
		itemSteelBeam,
		itemEncasedIndustrialBeam,
		itemModularFrame,
		itemHeavyModularFrame,
		itemFusedModularFrame,
		itemFicsiteTrigon,
		itemFabric,
		itemPlastic,
		itemRubber,
		itemRotor,
		itemStator,
		itemBattery,
		itemMotor,
		itemHeatSink,
		itemCoolingSystem,
		itemTurboMotor,
		itemWire,
		itemCable,
		itemQuickwire,
		itemCircuitBoard,
		itemAILimiter,
		itemHighSpeedConnector,
		itemReanimatedSAM,
		itemSAMFluctuator,
		itemComputer,
		itemSupercomputer,
		itemRadioControlUnit,
		itemCrystalOscillator,
		itemSuperpositionOscillator,
		itemDiamonds,
		itemTimeCrystal,
		itemDarkMatterCrystal,
		itemSingularityCell,
		itemNeuralQuantumProcessor,
		itemAlienPowerMatrix,
		itemEmptyCanister,
		itemEmptyFluidTank,
		itemPressureConversionCube,
		itemPackagedWater,
		itemPackagedAluminaSolution,
		itemPackagedSulfuricAcid,
		itemPackagedNitricAcid,
		itemPackagedNitrogenGas,
		itemLeaves,
		itemMycelia,
		itemFlowerPetals,
		itemWood,
		itemBiomass,
		itemCompactedCoal,
		itemPackagedOil,
		itemPackagedHeavyOilResidue,
		itemSolidBiofuel,
		itemPackagedFuel,
		itemPackagedLiquidBiofuel,
		itemPackagedTurbofuel,
		itemPackagedRocketFuel,
		itemPackagedIonizedFuel,
		itemUraniumFuelRod,
		itemPlutoniumFuelRod,
		itemBlackPowder,
		itemSmokelessPowder,
		itemGasFilter,
		itemColorCartridge,
		itemBeacon,
		itemIodineInfusedFilter,
		itemIronRebar,
		itemStunRebar,
		itemShatterRebar,
		itemExplosiveRebar,
		itemRifleAmmo,
		itemHomingRifleAmmo,
		itemTurboRifleAmmo,
		itemNobelisk,
		itemGasNobelisk,
		itemPulseNobelisk,
		itemClusterNobelisk,
		itemNukeNobelisk,
		itemElectromagneticControlRod,
		itemEncasedUraniumCell,
		itemNonFissileUranium,
		itemPlutoniumPellet,
		itemEncasedPlutoniumCell,
		itemFicsonium,
		itemFicsoniumFuelRod,
		itemUraniumWaste,
		itemPlutoniumWaste,
		itemBluePowerSlug,
		itemYellowPowerSlug,
		itemPurplePowerSlug,
		itemPowerShard,
		itemFICSITCoupon,
		itemSmartPlating,
		itemVersatileFramework,
		itemAutomatedWiring,
		itemModularEngine,
		itemAdaptiveControlUnit,
		itemAssemblyDirectorSystem,
		itemMagneticFieldGenerator,
		itemThermalPropulsionRocket,
		itemNuclearPasta,
		itemBiochemicalSculptor,
		itemBallisticWarpDrive,
		itemAIExpansionServer,
		itemPortableMiner,
		itemHatcherProtein,
		itemHogProtein,
		itemSpitterProtein,
		itemStingerProtein,
	)

	err = s.DB.Save(items).Error
	if err != nil {
		return err
	}

	recipes := []models.Recipe{
		{Slug: "adaptive_control_unit", Name: "Adaptive Control Unit", Duration: 600, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemAutomatedWiring, Qt: 50}, {Item: *itemCircuitBoard, Qt: 50}, {Item: *itemHeavyModularFrame, Qt: 10}, {Item: *itemComputer, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemAdaptiveControlUnit, Qt: 10}}},
		{Slug: "adhered_iron_plate", Name: "Adhered Iron Plate", Duration: 160, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronPlate, Qt: 30}, {Item: *itemRubber, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemReinforcedIronPlate, Qt: 10}}},
		{Slug: "ai_expansion_server", Name: "AI Expansion Server", Duration: 150, ProducedIn: buildingQuantumEncoder, Input: []models.RecipeInput{{Item: *itemAssemblyDirectorSystem, Qt: 10}, {Item: *itemNeuralQuantumProcessor, Qt: 10}, {Item: *itemSuperpositionOscillator, Qt: 10}, {Item: *itemExcitedPhotonicMatter, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemAIExpansionServer, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 250}}},
		{Slug: "ai_limiter", Name: "AI Limiter", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCopperSheet, Qt: 50}, {Item: *itemQuickwire, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemAILimiter, Qt: 10}}},
		{Slug: "alclad_aluminum_sheet", Name: "Alclad Aluminum Sheet", Duration: 60, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemAluminumIngot, Qt: 30}, {Item: *itemCopperIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAlcladAluminumSheet, Qt: 30}}},
		{Slug: "alclad_casing", Name: "Alclad Casing", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemAluminumIngot, Qt: 200}, {Item: *itemCopperIngot, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemAluminumCasing, Qt: 150}}},
		{Slug: "alien_dna_capsule", Name: "Alien DNA Capsule", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemAlienProtein, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAlienDNACapsule, Qt: 10}}},
		{Slug: "alien_power_matrix", Name: "Alien Power Matrix", Duration: 240, ProducedIn: buildingQuantumEncoder, Input: []models.RecipeInput{{Item: *itemSAMFluctuator, Qt: 50}, {Item: *itemPowerShard, Qt: 30}, {Item: *itemSuperpositionOscillator, Qt: 30}, {Item: *itemExcitedPhotonicMatter, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemAlienPowerMatrix, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 240}}},
		{Slug: "alumina_solution", Name: "Alumina Solution", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemBauxite, Qt: 120}, {Item: *itemWater, Qt: 180}}, Output: []models.RecipeOutput{{Item: *itemAluminaSolution, Qt: 120}, {Item: *itemSilica, Qt: 50}}},
		{Slug: "aluminum_beam", Name: "Aluminum Beam", Duration: 80, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemAluminumIngot, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSteelBeam, Qt: 30}}},
		{Slug: "aluminum_casing", Name: "Aluminum Casing", Duration: 20, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemAluminumIngot, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemAluminumCasing, Qt: 20}}},
		{Slug: "aluminum_ingot", Name: "Aluminum Ingot", Duration: 40, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemAluminumScrap, Qt: 60}, {Item: *itemSilica, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemAluminumIngot, Qt: 40}}},
		{Slug: "aluminum_rod", Name: "Aluminum Rod", Duration: 80, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemAluminumIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronRod, Qt: 70}}},
		{Slug: "aluminum_scrap", Name: "Aluminum Scrap", Duration: 10, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemAluminaSolution, Qt: 40}, {Item: *itemCoal, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemAluminumScrap, Qt: 60}, {Item: *itemWater, Qt: 20}}},
		{Slug: "assembly_director_system", Name: "Assembly Director System", Duration: 800, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemAdaptiveControlUnit, Qt: 20}, {Item: *itemSupercomputer, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemMagneticFieldGenerator, Qt: 10}}},
		{Slug: "automated_miner", Name: "Automated Miner", Duration: 600, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSteelPipe, Qt: 40}, {Item: *itemIronPlate, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemPortableMiner, Qt: 10}}},
		{Slug: "automated_speed_wiring", Name: "Automated Speed Wiring", Duration: 320, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemStator, Qt: 20}, {Item: *itemWire, Qt: 400}, {Item: *itemHighSpeedConnector, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAutomatedWiring, Qt: 40}}},
		{Slug: "automated_wiring", Name: "Automated Wiring", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemStator, Qt: 10}, {Item: *itemCable, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemAutomatedWiring, Qt: 10}}},
		{Slug: "ballistic_warp_drive", Name: "Ballistic Warp Drive", Duration: 600, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemThermalPropulsionRocket, Qt: 10}, {Item: *itemSingularityCell, Qt: 50}, {Item: *itemSuperpositionOscillator, Qt: 20}, {Item: *itemDarkMatterCrystal, Qt: 400}}, Output: []models.RecipeOutput{{Item: *itemBallisticWarpDrive, Qt: 10}}},
		{Slug: "basic_iron_ingot", Name: "Basic Iron Ingot", Duration: 120, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 50}, {Item: *itemLimestone, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemIronIngot, Qt: 100}}},
		{Slug: "battery", Name: "Battery", Duration: 30, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemSulfuricAcid, Qt: 25}, {Item: *itemAluminaSolution, Qt: 20}, {Item: *itemAluminumCasing, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemBattery, Qt: 10}, {Item: *itemWater, Qt: 15}}},
		{Slug: "bauxite_caterium", Name: "Bauxite (Caterium)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemCateriumOre, Qt: 150}}, Output: []models.RecipeOutput{{Item: *itemBauxite, Qt: 120}}},
		{Slug: "bauxite_copper", Name: "Bauxite (Copper)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemCopperOre, Qt: 180}}, Output: []models.RecipeOutput{{Item: *itemBauxite, Qt: 120}}},
		{Slug: "biochemical_sculptor", Name: "Biochemical Sculptor", Duration: 1200, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemMagneticFieldGenerator, Qt: 10}, {Item: *itemFicsiteTrigon, Qt: 800}, {Item: *itemWater, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemBiochemicalSculptor, Qt: 40}}},
		{Slug: "biocoal", Name: "Biocoal", Duration: 80, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemBiomass, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemCoal, Qt: 60}}},
		{Slug: "biomass_alien_protein", Name: "Biomass (Alien Protein)", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemAlienProtein, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemBiomass, Qt: 1000}}},
		{Slug: "biomass_leaves", Name: "Biomass (Leaves)", Duration: 50, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemLeaves, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemBiomass, Qt: 50}}},
		{Slug: "biomass_mycelia", Name: "Biomass (Mycelia)", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemMycelia, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemBiomass, Qt: 100}}},
		{Slug: "biomass_wood", Name: "Biomass (Wood)", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemWood, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemBiomass, Qt: 200}}},
		{Slug: "black_powder", Name: "Black Powder", Duration: 40, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCoal, Qt: 10}, {Item: *itemSulfur, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemBlackPowder, Qt: 20}}},
		{Slug: "bolted_frame", Name: "Bolted Frame", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemReinforcedIronPlate, Qt: 30}, {Item: *itemScrews, Qt: 560}}, Output: []models.RecipeOutput{{Item: *itemModularFrame, Qt: 20}}},
		{Slug: "bolted_iron_plate", Name: "Bolted Iron Plate", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronPlate, Qt: 180}, {Item: *itemScrews, Qt: 500}}, Output: []models.RecipeOutput{{Item: *itemReinforcedIronPlate, Qt: 30}}},
		{Slug: "cable", Name: "Cable", Duration: 20, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemWire, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCable, Qt: 10}}},
		{Slug: "cast_screws", Name: "Cast Screws", Duration: 240, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemScrews, Qt: 200}}},
		{Slug: "caterium_circuit_board", Name: "Caterium Circuit Board", Duration: 480, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemPlastic, Qt: 100}, {Item: *itemQuickwire, Qt: 300}}, Output: []models.RecipeOutput{{Item: *itemCircuitBoard, Qt: 70}}},
		{Slug: "caterium_computer", Name: "Caterium Computer", Duration: 160, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemCircuitBoard, Qt: 40}, {Item: *itemQuickwire, Qt: 140}, {Item: *itemRubber, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemComputer, Qt: 10}}},
		{Slug: "caterium_ingot", Name: "Caterium Ingot", Duration: 40, ProducedIn: buildingSmelter, Input: []models.RecipeInput{{Item: *itemCateriumOre, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemCateriumIngot, Qt: 10}}},
		{Slug: "caterium_ore_copper", Name: "Caterium Ore (Copper)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemCopperOre, Qt: 150}}, Output: []models.RecipeOutput{{Item: *itemCateriumOre, Qt: 120}}},
		{Slug: "caterium_ore_quartz", Name: "Caterium Ore (Quartz)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemRawQuartz, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemCateriumOre, Qt: 120}}},
		{Slug: "caterium_wire", Name: "Caterium Wire", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemCateriumIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemWire, Qt: 80}}},
		{Slug: "charcoal", Name: "Charcoal", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemWood, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemCoal, Qt: 100}}},
		{Slug: "cheap_silica", Name: "Cheap Silica", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemRawQuartz, Qt: 30}, {Item: *itemLimestone, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemSilica, Qt: 70}}},
		{Slug: "circuit_board", Name: "Circuit Board", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCopperSheet, Qt: 20}, {Item: *itemPlastic, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemCircuitBoard, Qt: 10}}},
		{Slug: "classic_battery", Name: "Classic Battery", Duration: 80, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemSulfur, Qt: 60}, {Item: *itemAlcladAluminumSheet, Qt: 70}, {Item: *itemPlastic, Qt: 80}, {Item: *itemWire, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemBattery, Qt: 40}}},
		{Slug: "cloudy_diamonds", Name: "Cloudy Diamonds", Duration: 30, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemCoal, Qt: 120}, {Item: *itemLimestone, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemDiamonds, Qt: 10}}},
		{Slug: "cluster_nobelisk", Name: "Cluster Nobelisk", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemNobelisk, Qt: 30}, {Item: *itemSmokelessPowder, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemClusterNobelisk, Qt: 10}}},
		{Slug: "coal_iron", Name: "Coal (Iron)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemIronOre, Qt: 180}}, Output: []models.RecipeOutput{{Item: *itemCoal, Qt: 120}}},
		{Slug: "coal_limestone", Name: "Coal (Limestone)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemLimestone, Qt: 360}}, Output: []models.RecipeOutput{{Item: *itemCoal, Qt: 120}}},
		{Slug: "coated_cable", Name: "Coated Cable", Duration: 80, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemWire, Qt: 50}, {Item: *itemHeavyOilResidue, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCable, Qt: 90}}},
		{Slug: "coated_iron_canister", Name: "Coated Iron Canister", Duration: 40, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronPlate, Qt: 20}, {Item: *itemCopperSheet, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemEmptyCanister, Qt: 40}}},
		{Slug: "coated_iron_plate", Name: "Coated Iron Plate", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 50}, {Item: *itemPlastic, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronPlate, Qt: 100}}},
		{Slug: "coke_steel_ingot", Name: "Coke Steel Ingot", Duration: 120, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 150}, {Item: *itemPetroleumCoke, Qt: 150}}, Output: []models.RecipeOutput{{Item: *itemSteelIngot, Qt: 200}}},
		{Slug: "compacted_coal", Name: "Compacted Coal", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCoal, Qt: 50}, {Item: *itemSulfur, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemCompactedCoal, Qt: 50}}},
		{Slug: "compacted_steel_ingot", Name: "Compacted Steel Ingot", Duration: 240, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 20}, {Item: *itemCompactedCoal, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemSteelIngot, Qt: 40}}},
		{Slug: "computer", Name: "Computer", Duration: 240, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemCircuitBoard, Qt: 40}, {Item: *itemCable, Qt: 80}, {Item: *itemPlastic, Qt: 160}}, Output: []models.RecipeOutput{{Item: *itemComputer, Qt: 10}}},
		{Slug: "concrete", Name: "Concrete", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemLimestone, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemConcrete, Qt: 10}}},
		{Slug: "cooling_device", Name: "Cooling Device", Duration: 240, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemHeatSink, Qt: 40}, {Item: *itemMotor, Qt: 10}, {Item: *itemNitrogenGas, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemCoolingSystem, Qt: 20}}},
		{Slug: "cooling_system", Name: "Cooling System", Duration: 100, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemHeatSink, Qt: 20}, {Item: *itemRubber, Qt: 20}, {Item: *itemWater, Qt: 50}, {Item: *itemNitrogenGas, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemCoolingSystem, Qt: 10}}},
		{Slug: "copper_alloy_ingot", Name: "Copper Alloy Ingot", Duration: 60, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemCopperOre, Qt: 50}, {Item: *itemIronOre, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemCopperIngot, Qt: 100}}},
		{Slug: "copper_ingot", Name: "Copper Ingot", Duration: 20, ProducedIn: buildingSmelter, Input: []models.RecipeInput{{Item: *itemCopperOre, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemCopperIngot, Qt: 10}}},
		{Slug: "copper_ore_quartz", Name: "Copper Ore (Quartz)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemRawQuartz, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemCopperOre, Qt: 120}}},
		{Slug: "copper_ore_sulfur", Name: "Copper Ore (Sulfur)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemSulfur, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemCopperOre, Qt: 120}}},
		{Slug: "copper_powder", Name: "Copper Powder", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemCopperIngot, Qt: 300}}, Output: []models.RecipeOutput{{Item: *itemCopperPowder, Qt: 50}}},
		{Slug: "copper_rotor", Name: "Copper Rotor", Duration: 160, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCopperSheet, Qt: 60}, {Item: *itemScrews, Qt: 520}}, Output: []models.RecipeOutput{{Item: *itemRotor, Qt: 30}}},
		{Slug: "copper_sheet", Name: "Copper Sheet", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemCopperIngot, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCopperSheet, Qt: 10}}},
		{Slug: "crystal_computer", Name: "Crystal Computer", Duration: 360, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCircuitBoard, Qt: 30}, {Item: *itemCrystalOscillator, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemComputer, Qt: 20}}},
		{Slug: "crystal_oscillator", Name: "Crystal Oscillator", Duration: 1200, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemQuartzCrystal, Qt: 360}, {Item: *itemCable, Qt: 280}, {Item: *itemReinforcedIronPlate, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemCrystalOscillator, Qt: 20}}},
		{Slug: "dark-ion_fuel", Name: "Dark-Ion Fuel", Duration: 30, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemPackagedRocketFuel, Qt: 120}, {Item: *itemDarkMatterCrystal, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemIonizedFuel, Qt: 100}, {Item: *itemCompactedCoal, Qt: 20}}},
		{Slug: "dark_matter_crystal", Name: "Dark Matter Crystal", Duration: 20, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemDiamonds, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemDarkMatterCrystal, Qt: 10}}},
		{Slug: "dark_matter_crystallization", Name: "Dark Matter Crystallization", Duration: 30, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemDarkMatterResidue, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemDarkMatterCrystal, Qt: 10}}},
		{Slug: "dark_matter_residue", Name: "Dark Matter Residue", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemDarkMatterResidue, Qt: 100}}},
		{Slug: "dark_matter_trap", Name: "Dark Matter Trap", Duration: 20, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemTimeCrystal, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemDarkMatterCrystal, Qt: 20}}},
		{Slug: "diamonds", Name: "Diamonds", Duration: 20, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemCoal, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemDiamonds, Qt: 10}}},
		{Slug: "diluted_fuel", Name: "Diluted Fuel", Duration: 60, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemHeavyOilResidue, Qt: 50}, {Item: *itemWater, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemFuel, Qt: 100}}},
		{Slug: "diluted_packaged_fuel", Name: "Diluted Packaged Fuel", Duration: 20, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemHeavyOilResidue, Qt: 10}, {Item: *itemPackagedWater, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedFuel, Qt: 20}}},
		{Slug: "distilled_silica", Name: "Distilled Silica", Duration: 60, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemDissolvedSilica, Qt: 120}, {Item: *itemLimestone, Qt: 50}, {Item: *itemWater, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemSilica, Qt: 270}, {Item: *itemWater, Qt: 80}}},
		{Slug: "electric_motor", Name: "Electric Motor", Duration: 160, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemElectromagneticControlRod, Qt: 10}, {Item: *itemRotor, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemMotor, Qt: 20}}},
		{Slug: "electrode_aluminum_scrap", Name: "Electrode Aluminum Scrap", Duration: 40, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemAluminaSolution, Qt: 120}, {Item: *itemPetroleumCoke, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemAluminumScrap, Qt: 200}, {Item: *itemWater, Qt: 70}}},
		{Slug: "electrode_circuit_board", Name: "Electrode Circuit Board", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemRubber, Qt: 40}, {Item: *itemPetroleumCoke, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemCircuitBoard, Qt: 10}}},
		{Slug: "electromagnetic_connection_rod", Name: "Electromagnetic Connection Rod", Duration: 150, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemStator, Qt: 20}, {Item: *itemHighSpeedConnector, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemElectromagneticControlRod, Qt: 20}}},
		{Slug: "electromagnetic_control_rod", Name: "Electromagnetic Control Rod", Duration: 300, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemStator, Qt: 30}, {Item: *itemAILimiter, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemElectromagneticControlRod, Qt: 20}}},
		{Slug: "empty_canister", Name: "Empty Canister", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemPlastic, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemEmptyCanister, Qt: 40}}},
		{Slug: "empty_fluid_tank", Name: "Empty Fluid Tank", Duration: 10, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemAluminumIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemEmptyFluidTank, Qt: 10}}},
		{Slug: "encased_industrial_beam", Name: "Encased Industrial Beam", Duration: 100, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSteelBeam, Qt: 30}, {Item: *itemConcrete, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemEncasedIndustrialBeam, Qt: 10}}},
		{Slug: "encased_industrial_pipe", Name: "Encased Industrial Pipe", Duration: 150, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSteelPipe, Qt: 60}, {Item: *itemConcrete, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemEncasedIndustrialBeam, Qt: 10}}},
		{Slug: "encased_plutonium_cell", Name: "Encased Plutonium Cell", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemPlutoniumPellet, Qt: 20}, {Item: *itemConcrete, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemEncasedPlutoniumCell, Qt: 10}}},
		{Slug: "encased_uranium_cell", Name: "Encased Uranium Cell", Duration: 120, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemUranium, Qt: 100}, {Item: *itemConcrete, Qt: 30}, {Item: *itemSulfuricAcid, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemEncasedUraniumCell, Qt: 50}, {Item: *itemSulfuricAcid, Qt: 20}}},
		{Slug: "excited_photonic_matter", Name: "Excited Photonic Matter", Duration: 30, ProducedIn: buildingConverter, Input: []models.RecipeInput{}, Output: []models.RecipeOutput{{Item: *itemExcitedPhotonicMatter, Qt: 100}}},
		{Slug: "explosive_rebar", Name: "Explosive Rebar", Duration: 120, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemIronRebar, Qt: 20}, {Item: *itemSmokelessPowder, Qt: 20}, {Item: *itemSteelPipe, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemExplosiveRebar, Qt: 10}}},
		{Slug: "fabric", Name: "Fabric", Duration: 40, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemMycelia, Qt: 10}, {Item: *itemBiomass, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemFabric, Qt: 10}}},
		{Slug: "fertile_uranium", Name: "Fertile Uranium", Duration: 120, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemUranium, Qt: 50}, {Item: *itemUraniumWaste, Qt: 50}, {Item: *itemNitricAcid, Qt: 30}, {Item: *itemSulfuricAcid, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemNonFissileUranium, Qt: 200}, {Item: *itemWater, Qt: 80}}},
		{Slug: "ficsite_ingot_aluminum", Name: "Ficsite Ingot (Aluminum)", Duration: 20, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 20}, {Item: *itemAluminumIngot, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemFicsiteIngot, Qt: 10}}},
		{Slug: "ficsite_ingot_caterium", Name: "Ficsite Ingot (Caterium)", Duration: 40, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 30}, {Item: *itemCateriumIngot, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemFicsiteIngot, Qt: 10}}},
		{Slug: "ficsite_ingot_iron", Name: "Ficsite Ingot (Iron)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 40}, {Item: *itemIronIngot, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemFicsiteIngot, Qt: 10}}},
		{Slug: "ficsite_trigon", Name: "Ficsite Trigon", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemFicsiteIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemFicsiteTrigon, Qt: 30}}},
		{Slug: "ficsonium", Name: "Ficsonium", Duration: 60, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemPlutoniumWaste, Qt: 10}, {Item: *itemSingularityCell, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemFicsonium, Qt: 10}}},
		{Slug: "ficsonium_fuel_rod", Name: "Ficsonium Fuel Rod", Duration: 240, ProducedIn: buildingQuantumEncoder, Input: []models.RecipeInput{{Item: *itemFicsonium, Qt: 20}, {Item: *itemElectromagneticControlRod, Qt: 20}, {Item: *itemFicsiteTrigon, Qt: 400}, {Item: *itemExcitedPhotonicMatter, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemFicsoniumFuelRod, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 200}}},
		{Slug: "fine_black_powder", Name: "Fine Black Powder", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSulfur, Qt: 10}, {Item: *itemCompactedCoal, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemBlackPowder, Qt: 60}}},
		{Slug: "fine_concrete", Name: "Fine Concrete", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSilica, Qt: 30}, {Item: *itemLimestone, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemConcrete, Qt: 100}}},
		{Slug: "flexible_framework", Name: "Flexible Framework", Duration: 160, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemModularFrame, Qt: 10}, {Item: *itemSteelBeam, Qt: 60}, {Item: *itemRubber, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemVersatileFramework, Qt: 20}}},
		{Slug: "fuel", Name: "Fuel", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemFuel, Qt: 40}, {Item: *itemPolymerResin, Qt: 30}}},
		{Slug: "fused_modular_frame", Name: "Fused Modular Frame", Duration: 400, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemHeavyModularFrame, Qt: 10}, {Item: *itemAluminumCasing, Qt: 500}, {Item: *itemNitrogenGas, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemFusedModularFrame, Qt: 10}}},
		{Slug: "fused_quartz_crystal", Name: "Fused Quartz Crystal", Duration: 200, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemRawQuartz, Qt: 250}, {Item: *itemCoal, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemQuartzCrystal, Qt: 180}}},
		{Slug: "fused_quickwire", Name: "Fused Quickwire", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCateriumIngot, Qt: 10}, {Item: *itemCopperIngot, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemQuickwire, Qt: 120}}},
		{Slug: "fused_wire", Name: "Fused Wire", Duration: 200, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCopperIngot, Qt: 40}, {Item: *itemCateriumIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemWire, Qt: 300}}},
		{Slug: "gas_filter", Name: "Gas Filter", Duration: 80, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemFabric, Qt: 20}, {Item: *itemCoal, Qt: 40}, {Item: *itemIronPlate, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemGasFilter, Qt: 10}}},
		{Slug: "gas_nobelisk", Name: "Gas Nobelisk", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemNobelisk, Qt: 10}, {Item: *itemBiomass, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemGasNobelisk, Qt: 10}}},
		{Slug: "hatcher_protein", Name: "Hatcher Protein", Duration: 30, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemHatcherProtein, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAlienProtein, Qt: 10}}},
		{Slug: "heat-fused_frame", Name: "Heat-Fused Frame", Duration: 200, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemHeavyModularFrame, Qt: 10}, {Item: *itemAluminumIngot, Qt: 500}, {Item: *itemNitricAcid, Qt: 80}, {Item: *itemFuel, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemFusedModularFrame, Qt: 10}}},
		{Slug: "heat_exchanger", Name: "Heat Exchanger", Duration: 60, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemAluminumCasing, Qt: 30}, {Item: *itemRubber, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemHeatSink, Qt: 10}}},
		{Slug: "heat_sink", Name: "Heat Sink", Duration: 80, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemAlcladAluminumSheet, Qt: 50}, {Item: *itemCopperSheet, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemHeatSink, Qt: 10}}},
		{Slug: "heavy_encased_frame", Name: "Heavy Encased Frame", Duration: 640, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemModularFrame, Qt: 80}, {Item: *itemEncasedIndustrialBeam, Qt: 100}, {Item: *itemSteelPipe, Qt: 360}, {Item: *itemConcrete, Qt: 220}}, Output: []models.RecipeOutput{{Item: *itemHeavyModularFrame, Qt: 30}}},
		{Slug: "heavy_flexible_frame", Name: "Heavy Flexible Frame", Duration: 160, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemModularFrame, Qt: 50}, {Item: *itemEncasedIndustrialBeam, Qt: 30}, {Item: *itemRubber, Qt: 200}, {Item: *itemScrews, Qt: 1040}}, Output: []models.RecipeOutput{{Item: *itemHeavyModularFrame, Qt: 10}}},
		{Slug: "heavy_modular_frame", Name: "Heavy Modular Frame", Duration: 300, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemModularFrame, Qt: 50}, {Item: *itemSteelPipe, Qt: 200}, {Item: *itemEncasedIndustrialBeam, Qt: 50}, {Item: *itemScrews, Qt: 1200}}, Output: []models.RecipeOutput{{Item: *itemHeavyModularFrame, Qt: 10}}},
		{Slug: "heavy_oil_residue", Name: "Heavy Oil Residue", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemHeavyOilResidue, Qt: 40}, {Item: *itemPolymerResin, Qt: 20}}},
		{Slug: "high-speed_connector", Name: "High-Speed Connector", Duration: 160, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemQuickwire, Qt: 560}, {Item: *itemCable, Qt: 100}, {Item: *itemCircuitBoard, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemHighSpeedConnector, Qt: 10}}},
		{Slug: "hog_protein", Name: "Hog Protein", Duration: 30, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemHogProtein, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAlienProtein, Qt: 10}}},
		{Slug: "homing_rifle_ammo", Name: "Homing Rifle Ammo", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemRifleAmmo, Qt: 200}, {Item: *itemHighSpeedConnector, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemHomingRifleAmmo, Qt: 100}}},
		{Slug: "infused_uranium_cell", Name: "Infused Uranium Cell", Duration: 120, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemUranium, Qt: 50}, {Item: *itemSilica, Qt: 30}, {Item: *itemSulfur, Qt: 50}, {Item: *itemQuickwire, Qt: 150}}, Output: []models.RecipeOutput{{Item: *itemEncasedUraniumCell, Qt: 40}}},
		{Slug: "instant_plutonium_cell", Name: "Instant Plutonium Cell", Duration: 1200, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemNonFissileUranium, Qt: 1500}, {Item: *itemAluminumCasing, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemEncasedPlutoniumCell, Qt: 200}}},
		{Slug: "instant_scrap", Name: "Instant Scrap", Duration: 60, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemBauxite, Qt: 150}, {Item: *itemCoal, Qt: 100}, {Item: *itemSulfuricAcid, Qt: 50}, {Item: *itemWater, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemAluminumScrap, Qt: 300}, {Item: *itemWater, Qt: 50}}},
		{Slug: "insulated_cable", Name: "Insulated Cable", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemWire, Qt: 90}, {Item: *itemRubber, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemCable, Qt: 200}}},
		{Slug: "insulated_crystal_oscillator", Name: "Insulated Crystal Oscillator", Duration: 320, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemQuartzCrystal, Qt: 100}, {Item: *itemRubber, Qt: 70}, {Item: *itemAILimiter, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemCrystalOscillator, Qt: 10}}},
		{Slug: "iodine-infused_filter", Name: "Iodine-Infused Filter", Duration: 160, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemGasFilter, Qt: 10}, {Item: *itemQuickwire, Qt: 80}, {Item: *itemAluminumCasing, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIodineInfusedFilter, Qt: 10}}},
		{Slug: "ionized_fuel", Name: "Ionized Fuel", Duration: 240, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemRocketFuel, Qt: 160}, {Item: *itemPowerShard, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIonizedFuel, Qt: 160}, {Item: *itemCompactedCoal, Qt: 20}}},
		{Slug: "iron_alloy_ingot", Name: "Iron Alloy Ingot", Duration: 120, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 80}, {Item: *itemCopperOre, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemIronIngot, Qt: 150}}},
		{Slug: "iron_ingot", Name: "Iron Ingot", Duration: 20, ProducedIn: buildingSmelter, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronIngot, Qt: 10}}},
		{Slug: "iron_ore_limestone", Name: "Iron Ore (Limestone)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemLimestone, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemIronOre, Qt: 120}}},
		{Slug: "iron_pipe", Name: "Iron Pipe", Duration: 120, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemSteelPipe, Qt: 50}}},
		{Slug: "iron_plate", Name: "Iron Plate", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemIronPlate, Qt: 20}}},
		{Slug: "iron_rebar", Name: "Iron Rebar", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronRod, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronRebar, Qt: 10}}},
		{Slug: "iron_rod", Name: "Iron Rod", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronRod, Qt: 10}}},
		{Slug: "iron_wire", Name: "Iron Wire", Duration: 240, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemWire, Qt: 90}}},
		{Slug: "leached_caterium_ingot", Name: "Leached Caterium Ingot", Duration: 100, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCateriumOre, Qt: 90}, {Item: *itemSulfuricAcid, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemCateriumIngot, Qt: 60}}},
		{Slug: "leached_copper_ingot", Name: "Leached Copper Ingot", Duration: 120, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCopperOre, Qt: 90}, {Item: *itemSulfuricAcid, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemCopperIngot, Qt: 220}}},
		{Slug: "leached_iron_ingot", Name: "Leached Iron Ingot", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 50}, {Item: *itemSulfuricAcid, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronIngot, Qt: 100}}},
		{Slug: "limestone_sulfur", Name: "Limestone (Sulfur)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemSulfur, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemLimestone, Qt: 120}}},
		{Slug: "liquid_biofuel", Name: "Liquid Biofuel", Duration: 40, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemSolidBiofuel, Qt: 60}, {Item: *itemWater, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemLiquidBiofuel, Qt: 40}}},
		{Slug: "magnetic_field_generator", Name: "Magnetic Field Generator", Duration: 1200, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemVersatileFramework, Qt: 50}, {Item: *itemElectromagneticControlRod, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemAssemblyDirectorSystem, Qt: 20}}},
		{Slug: "modular_engine", Name: "Modular Engine", Duration: 600, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemMotor, Qt: 20}, {Item: *itemRubber, Qt: 150}, {Item: *itemSmartPlating, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemModularEngine, Qt: 10}}},
		{Slug: "modular_frame", Name: "Modular Frame", Duration: 600, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemReinforcedIronPlate, Qt: 30}, {Item: *itemIronRod, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemModularFrame, Qt: 20}}},
		{Slug: "molded_beam", Name: "Molded Beam", Duration: 120, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemSteelIngot, Qt: 240}, {Item: *itemConcrete, Qt: 160}}, Output: []models.RecipeOutput{{Item: *itemSteelBeam, Qt: 90}}},
		{Slug: "molded_steel_pipe", Name: "Molded Steel Pipe", Duration: 60, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemSteelIngot, Qt: 50}, {Item: *itemConcrete, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSteelPipe, Qt: 50}}},
		{Slug: "motor", Name: "Motor", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemRotor, Qt: 20}, {Item: *itemStator, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemMotor, Qt: 10}}},
		{Slug: "neural-quantum_processor", Name: "Neural-Quantum Processor", Duration: 200, ProducedIn: buildingQuantumEncoder, Input: []models.RecipeInput{{Item: *itemTimeCrystal, Qt: 50}, {Item: *itemSupercomputer, Qt: 10}, {Item: *itemFicsiteTrigon, Qt: 150}, {Item: *itemExcitedPhotonicMatter, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemNeuralQuantumProcessor, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 250}}},
		{Slug: "nitric_acid", Name: "Nitric Acid", Duration: 60, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemNitrogenGas, Qt: 120}, {Item: *itemWater, Qt: 30}, {Item: *itemIronPlate, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemNitricAcid, Qt: 30}}},
		{Slug: "nitro_rocket_fuel", Name: "Nitro Rocket Fuel", Duration: 24, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemFuel, Qt: 40}, {Item: *itemNitrogenGas, Qt: 30}, {Item: *itemSulfur, Qt: 40}, {Item: *itemCoal, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemRocketFuel, Qt: 60}, {Item: *itemCompactedCoal, Qt: 10}}},
		{Slug: "nitrogen_gas_bauxite", Name: "Nitrogen Gas (Bauxite)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemBauxite, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemNitrogenGas, Qt: 120}}},
		{Slug: "nitrogen_gas_caterium", Name: "Nitrogen Gas (Caterium)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemCateriumOre, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemNitrogenGas, Qt: 120}}},
		{Slug: "nobelisk", Name: "Nobelisk", Duration: 60, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemBlackPowder, Qt: 20}, {Item: *itemSteelPipe, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemNobelisk, Qt: 10}}},
		{Slug: "non-fissile_uranium", Name: "Non-Fissile Uranium", Duration: 240, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemUraniumWaste, Qt: 150}, {Item: *itemSilica, Qt: 100}, {Item: *itemNitricAcid, Qt: 60}, {Item: *itemSulfuricAcid, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemNonFissileUranium, Qt: 200}, {Item: *itemWater, Qt: 60}}},
		{Slug: "nuclear_pasta", Name: "Nuclear Pasta", Duration: 1200, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemCopperPowder, Qt: 2000}, {Item: *itemPressureConversionCube, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemNuclearPasta, Qt: 10}}},
		{Slug: "nuke_nobelisk", Name: "Nuke Nobelisk", Duration: 1200, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemNobelisk, Qt: 50}, {Item: *itemEncasedUraniumCell, Qt: 200}, {Item: *itemSmokelessPowder, Qt: 100}, {Item: *itemAILimiter, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemNukeNobelisk, Qt: 10}}},
		{Slug: "oc_supercomputer", Name: "OC Supercomputer", Duration: 200, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemRadioControlUnit, Qt: 20}, {Item: *itemCoolingSystem, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemSupercomputer, Qt: 10}}},
		{Slug: "oil-based_diamonds", Name: "Oil-Based Diamonds", Duration: 30, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemDiamonds, Qt: 20}}},
		{Slug: "packaged_alumina_solution", Name: "Packaged Alumina Solution", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemAluminaSolution, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedAluminaSolution, Qt: 20}}},
		{Slug: "packaged_fuel", Name: "Packaged Fuel", Duration: 30, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemFuel, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedFuel, Qt: 20}}},
		{Slug: "packaged_heavy_oil_residue", Name: "Packaged Heavy Oil Residue", Duration: 40, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemHeavyOilResidue, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedHeavyOilResidue, Qt: 20}}},
		{Slug: "packaged_ionized_fuel", Name: "Packaged Ionized Fuel", Duration: 30, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemIonizedFuel, Qt: 40}, {Item: *itemEmptyFluidTank, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedIonizedFuel, Qt: 20}}},
		{Slug: "packaged_liquid_biofuel", Name: "Packaged Liquid Biofuel", Duration: 30, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemLiquidBiofuel, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedLiquidBiofuel, Qt: 20}}},
		{Slug: "packaged_nitric_acid", Name: "Packaged Nitric Acid", Duration: 20, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemNitricAcid, Qt: 10}, {Item: *itemEmptyFluidTank, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPackagedNitricAcid, Qt: 10}}},
		{Slug: "packaged_nitrogen_gas", Name: "Packaged Nitrogen Gas", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemNitrogenGas, Qt: 40}, {Item: *itemEmptyFluidTank, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPackagedNitrogenGas, Qt: 10}}},
		{Slug: "packaged_oil", Name: "Packaged Oil", Duration: 40, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedOil, Qt: 20}}},
		{Slug: "packaged_rocket_fuel", Name: "Packaged Rocket Fuel", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemRocketFuel, Qt: 20}, {Item: *itemEmptyFluidTank, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPackagedRocketFuel, Qt: 10}}},
		{Slug: "packaged_sulfuric_acid", Name: "Packaged Sulfuric Acid", Duration: 30, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemSulfuricAcid, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedSulfuricAcid, Qt: 20}}},
		{Slug: "packaged_turbofuel", Name: "Packaged Turbofuel", Duration: 60, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemTurbofuel, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedTurbofuel, Qt: 20}}},
		{Slug: "packaged_water", Name: "Packaged Water", Duration: 20, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemWater, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPackagedWater, Qt: 20}}},
		{Slug: "petroleum_coke", Name: "Petroleum Coke", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemHeavyOilResidue, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemPetroleumCoke, Qt: 120}}},
		{Slug: "petroleum_diamonds", Name: "Petroleum Diamonds", Duration: 20, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemPetroleumCoke, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemDiamonds, Qt: 10}}},
		{Slug: "pink_diamonds", Name: "Pink Diamonds", Duration: 40, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemCoal, Qt: 80}, {Item: *itemQuartzCrystal, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemDiamonds, Qt: 10}}},
		{Slug: "plastic", Name: "Plastic", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemPlastic, Qt: 20}, {Item: *itemHeavyOilResidue, Qt: 10}}},
		{Slug: "plastic_ai_limiter", Name: "Plastic AI Limiter", Duration: 150, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemQuickwire, Qt: 300}, {Item: *itemPlastic, Qt: 70}}, Output: []models.RecipeOutput{{Item: *itemAILimiter, Qt: 20}}},
		{Slug: "plastic_smart_plating", Name: "Plastic Smart Plating", Duration: 240, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemReinforcedIronPlate, Qt: 10}, {Item: *itemRotor, Qt: 10}, {Item: *itemPlastic, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSmartPlating, Qt: 20}}},
		{Slug: "plutonium_fuel_rod", Name: "Plutonium Fuel Rod", Duration: 2400, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemEncasedPlutoniumCell, Qt: 300}, {Item: *itemSteelBeam, Qt: 180}, {Item: *itemElectromagneticControlRod, Qt: 60}, {Item: *itemHeatSink, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemPlutoniumFuelRod, Qt: 10}}},
		{Slug: "plutonium_fuel_rod_burning", Name: "Plutonium Fuel Rod (burning)", Duration: 6000, ProducedIn: buildingNuclearGenerator, Input: []models.RecipeInput{{Item: *itemPlutoniumFuelRod, Qt: 10}, {Item: *itemWater, Qt: 24000}}, Output: []models.RecipeOutput{{Item: *itemPlutoniumWaste, Qt: 100}}},
		{Slug: "plutonium_fuel_unit", Name: "Plutonium Fuel Unit", Duration: 1200, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemEncasedPlutoniumCell, Qt: 200}, {Item: *itemPressureConversionCube, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPlutoniumFuelRod, Qt: 10}}},
		{Slug: "plutonium_pellet", Name: "Plutonium Pellet", Duration: 600, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemNonFissileUranium, Qt: 1000}, {Item: *itemUraniumWaste, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemPlutoniumPellet, Qt: 300}}},
		{Slug: "polyester_fabric", Name: "Polyester Fabric", Duration: 20, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemPolymerResin, Qt: 10}, {Item: *itemWater, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemFabric, Qt: 10}}},
		{Slug: "polymer_resin", Name: "Polymer Resin", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemPolymerResin, Qt: 130}, {Item: *itemHeavyOilResidue, Qt: 20}}},
		{Slug: "power_shard_1", Name: "Power Shard (1)", Duration: 80, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemBluePowerSlug, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPowerShard, Qt: 10}}},
		{Slug: "power_shard_2", Name: "Power Shard (2)", Duration: 120, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemYellowPowerSlug, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPowerShard, Qt: 20}}},
		{Slug: "power_shard_5", Name: "Power Shard (5)", Duration: 240, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemPurplePowerSlug, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPowerShard, Qt: 50}}},
		{Slug: "pressure_conversion_cube", Name: "Pressure Conversion Cube", Duration: 600, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemFusedModularFrame, Qt: 10}, {Item: *itemRadioControlUnit, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPressureConversionCube, Qt: 10}}},
		{Slug: "pulse_nobelisk", Name: "Pulse Nobelisk", Duration: 600, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemNobelisk, Qt: 50}, {Item: *itemCrystalOscillator, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemPulseNobelisk, Qt: 50}}},
		{Slug: "pure_aluminum_ingot", Name: "Pure Aluminum Ingot", Duration: 20, ProducedIn: buildingSmelter, Input: []models.RecipeInput{{Item: *itemAluminumScrap, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemAluminumIngot, Qt: 10}}},
		{Slug: "pure_caterium_ingot", Name: "Pure Caterium Ingot", Duration: 50, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCateriumOre, Qt: 20}, {Item: *itemWater, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCateriumIngot, Qt: 10}}},
		{Slug: "pure_copper_ingot", Name: "Pure Copper Ingot", Duration: 240, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCopperOre, Qt: 60}, {Item: *itemWater, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemCopperIngot, Qt: 150}}},
		{Slug: "pure_iron_ingot", Name: "Pure Iron Ingot", Duration: 120, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 70}, {Item: *itemWater, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemIronIngot, Qt: 130}}},
		{Slug: "pure_quartz_crystal", Name: "Pure Quartz Crystal", Duration: 80, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemRawQuartz, Qt: 90}, {Item: *itemWater, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemQuartzCrystal, Qt: 70}}},
		{Slug: "quartz_crystal", Name: "Quartz Crystal", Duration: 80, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemRawQuartz, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemQuartzCrystal, Qt: 30}}},
		{Slug: "quartz_purification", Name: "Quartz Purification", Duration: 120, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemRawQuartz, Qt: 240}, {Item: *itemNitricAcid, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemQuartzCrystal, Qt: 150}, {Item: *itemDissolvedSilica, Qt: 120}}},
		{Slug: "quickwire", Name: "Quickwire", Duration: 50, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemCateriumIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemQuickwire, Qt: 50}}},
		{Slug: "quickwire_cable", Name: "Quickwire Cable", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemQuickwire, Qt: 30}, {Item: *itemRubber, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCable, Qt: 110}}},
		{Slug: "quickwire_stator", Name: "Quickwire Stator", Duration: 150, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSteelPipe, Qt: 40}, {Item: *itemQuickwire, Qt: 150}}, Output: []models.RecipeOutput{{Item: *itemStator, Qt: 20}}},
		{Slug: "radio_connection_unit", Name: "Radio Connection Unit", Duration: 160, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemHeatSink, Qt: 40}, {Item: *itemHighSpeedConnector, Qt: 20}, {Item: *itemQuartzCrystal, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemRadioControlUnit, Qt: 10}}},
		{Slug: "radio_control_system", Name: "Radio Control System", Duration: 400, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemCrystalOscillator, Qt: 10}, {Item: *itemCircuitBoard, Qt: 100}, {Item: *itemAluminumCasing, Qt: 600}, {Item: *itemRubber, Qt: 300}}, Output: []models.RecipeOutput{{Item: *itemRadioControlUnit, Qt: 30}}},
		{Slug: "radio_control_unit", Name: "Radio Control Unit", Duration: 480, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemAluminumCasing, Qt: 320}, {Item: *itemCrystalOscillator, Qt: 10}, {Item: *itemComputer, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemRadioControlUnit, Qt: 20}}},
		{Slug: "raw_quartz_bauxite", Name: "Raw Quartz (Bauxite)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemBauxite, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemRawQuartz, Qt: 120}}},
		{Slug: "raw_quartz_coal", Name: "Raw Quartz (Coal)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemCoal, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemRawQuartz, Qt: 120}}},
		{Slug: "reanimated_sam", Name: "Reanimated SAM", Duration: 20, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSAM, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemReanimatedSAM, Qt: 10}}},
		{Slug: "recycled_plastic", Name: "Recycled Plastic", Duration: 120, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemRubber, Qt: 60}, {Item: *itemFuel, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemPlastic, Qt: 120}}},
		{Slug: "recycled_rubber", Name: "Recycled Rubber", Duration: 120, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemPlastic, Qt: 60}, {Item: *itemFuel, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemRubber, Qt: 120}}},
		{Slug: "reinforced_iron_plate", Name: "Reinforced Iron Plate", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronPlate, Qt: 60}, {Item: *itemScrews, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemReinforcedIronPlate, Qt: 10}}},
		{Slug: "residual_fuel", Name: "Residual Fuel", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemHeavyOilResidue, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemFuel, Qt: 40}}},
		{Slug: "residual_plastic", Name: "Residual Plastic", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemPolymerResin, Qt: 60}, {Item: *itemWater, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemPlastic, Qt: 20}}},
		{Slug: "residual_rubber", Name: "Residual Rubber", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemPolymerResin, Qt: 40}, {Item: *itemWater, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemRubber, Qt: 20}}},
		{Slug: "rifle_ammo", Name: "Rifle Ammo", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCopperSheet, Qt: 30}, {Item: *itemSmokelessPowder, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemRifleAmmo, Qt: 150}}},
		{Slug: "rigor_motor", Name: "Rigor Motor", Duration: 480, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemRotor, Qt: 30}, {Item: *itemStator, Qt: 30}, {Item: *itemCrystalOscillator, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemMotor, Qt: 60}}},
		{Slug: "rocket_fuel", Name: "Rocket Fuel", Duration: 60, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemTurbofuel, Qt: 60}, {Item: *itemNitricAcid, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemRocketFuel, Qt: 100}, {Item: *itemCompactedCoal, Qt: 10}}},
		{Slug: "rotor", Name: "Rotor", Duration: 150, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronRod, Qt: 50}, {Item: *itemScrews, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemRotor, Qt: 10}}},
		{Slug: "rubber", Name: "Rubber", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCrudeOil, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemRubber, Qt: 20}, {Item: *itemHeavyOilResidue, Qt: 20}}},
		{Slug: "rubber_concrete", Name: "Rubber Concrete", Duration: 60, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemLimestone, Qt: 100}, {Item: *itemRubber, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemConcrete, Qt: 90}}},
		{Slug: "sam_fluctuator", Name: "SAM Fluctuator", Duration: 60, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 60}, {Item: *itemWire, Qt: 50}, {Item: *itemSteelPipe, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSAMFluctuator, Qt: 10}}},
		{Slug: "screws", Name: "Screws", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemIronRod, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemScrews, Qt: 40}}},
		{Slug: "shatter_rebar", Name: "Shatter Rebar", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronRebar, Qt: 20}, {Item: *itemQuartzCrystal, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemShatterRebar, Qt: 10}}},
		{Slug: "silica", Name: "Silica", Duration: 80, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemRawQuartz, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSilica, Qt: 50}}},
		{Slug: "silicon_circuit_board", Name: "Silicon Circuit Board", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemCopperSheet, Qt: 110}, {Item: *itemSilica, Qt: 110}}, Output: []models.RecipeOutput{{Item: *itemCircuitBoard, Qt: 50}}},
		{Slug: "silicon_high-speed_connector", Name: "Silicon High-Speed Connector", Duration: 400, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemQuickwire, Qt: 600}, {Item: *itemSilica, Qt: 250}, {Item: *itemCircuitBoard, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemHighSpeedConnector, Qt: 20}}},
		{Slug: "singularity_cell", Name: "Singularity Cell", Duration: 600, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemNuclearPasta, Qt: 10}, {Item: *itemDarkMatterCrystal, Qt: 200}, {Item: *itemIronPlate, Qt: 1000}, {Item: *itemConcrete, Qt: 2000}}, Output: []models.RecipeOutput{{Item: *itemSingularityCell, Qt: 100}}},
		{Slug: "sloppy_alumina", Name: "Sloppy Alumina", Duration: 30, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemBauxite, Qt: 100}, {Item: *itemWater, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemAluminaSolution, Qt: 120}}},
		{Slug: "smart_plating", Name: "Smart Plating", Duration: 300, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemReinforcedIronPlate, Qt: 10}, {Item: *itemRotor, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemSmartPlating, Qt: 10}}},
		{Slug: "smokeless_powder", Name: "Smokeless Powder", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemBlackPowder, Qt: 20}, {Item: *itemHeavyOilResidue, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemSmokelessPowder, Qt: 20}}},
		{Slug: "solid_biofuel", Name: "Solid Biofuel", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemBiomass, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemSolidBiofuel, Qt: 40}}},
		{Slug: "solid_steel_ingot", Name: "Solid Steel Ingot", Duration: 30, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 20}, {Item: *itemCoal, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemSteelIngot, Qt: 30}}},
		{Slug: "spitter_protein", Name: "Spitter Protein", Duration: 30, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSpitterProtein, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAlienProtein, Qt: 10}}},
		{Slug: "stator", Name: "Stator", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSteelPipe, Qt: 30}, {Item: *itemWire, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemStator, Qt: 10}}},
		{Slug: "steamed_copper_sheet", Name: "Steamed Copper Sheet", Duration: 80, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemCopperIngot, Qt: 30}, {Item: *itemWater, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemCopperSheet, Qt: 30}}},
		{Slug: "steel_beam", Name: "Steel Beam", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSteelIngot, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemSteelBeam, Qt: 10}}},
		{Slug: "steel_canister", Name: "Steel Canister", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSteelIngot, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemEmptyCanister, Qt: 40}}},
		{Slug: "steel_cast_plate", Name: "Steel Cast Plate", Duration: 40, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronIngot, Qt: 10}, {Item: *itemSteelIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronPlate, Qt: 30}}},
		{Slug: "steel_ingot", Name: "Steel Ingot", Duration: 40, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemIronOre, Qt: 30}, {Item: *itemCoal, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSteelIngot, Qt: 30}}},
		{Slug: "steel_pipe", Name: "Steel Pipe", Duration: 60, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSteelIngot, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemSteelPipe, Qt: 20}}},
		{Slug: "steel_rod", Name: "Steel Rod", Duration: 50, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSteelIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemIronRod, Qt: 40}}},
		{Slug: "steel_rotor", Name: "Steel Rotor", Duration: 120, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemSteelPipe, Qt: 20}, {Item: *itemWire, Qt: 60}}, Output: []models.RecipeOutput{{Item: *itemRotor, Qt: 10}}},
		{Slug: "steel_screws", Name: "Steel Screws", Duration: 120, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemSteelBeam, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemScrews, Qt: 520}}},
		{Slug: "steeled_frame", Name: "Steeled Frame", Duration: 600, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemReinforcedIronPlate, Qt: 20}, {Item: *itemSteelPipe, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemModularFrame, Qt: 30}}},
		{Slug: "stinger_protein", Name: "Stinger Protein", Duration: 30, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemStingerProtein, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemAlienProtein, Qt: 10}}},
		{Slug: "stitched_iron_plate", Name: "Stitched Iron Plate", Duration: 320, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronPlate, Qt: 100}, {Item: *itemWire, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemReinforcedIronPlate, Qt: 30}}},
		{Slug: "stun_rebar", Name: "Stun Rebar", Duration: 60, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemIronRebar, Qt: 10}, {Item: *itemQuickwire, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemStunRebar, Qt: 10}}},
		{Slug: "sulfur_coal", Name: "Sulfur (Coal)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemCoal, Qt: 200}}, Output: []models.RecipeOutput{{Item: *itemSulfur, Qt: 120}}},
		{Slug: "sulfur_iron", Name: "Sulfur (Iron)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemIronOre, Qt: 300}}, Output: []models.RecipeOutput{{Item: *itemSulfur, Qt: 120}}},
		{Slug: "sulfuric_acid", Name: "Sulfuric Acid", Duration: 60, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemSulfur, Qt: 50}, {Item: *itemWater, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemSulfuricAcid, Qt: 50}}},
		{Slug: "super-state_computer", Name: "Super-State Computer", Duration: 250, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemComputer, Qt: 30}, {Item: *itemElectromagneticControlRod, Qt: 10}, {Item: *itemBattery, Qt: 100}, {Item: *itemWire, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemSupercomputer, Qt: 10}}},
		{Slug: "supercomputer", Name: "Supercomputer", Duration: 320, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemComputer, Qt: 40}, {Item: *itemAILimiter, Qt: 20}, {Item: *itemHighSpeedConnector, Qt: 30}, {Item: *itemPlastic, Qt: 280}}, Output: []models.RecipeOutput{{Item: *itemSupercomputer, Qt: 10}}},
		{Slug: "superposition_oscillator", Name: "Superposition Oscillator", Duration: 120, ProducedIn: buildingQuantumEncoder, Input: []models.RecipeInput{{Item: *itemDarkMatterCrystal, Qt: 60}, {Item: *itemCrystalOscillator, Qt: 10}, {Item: *itemAlcladAluminumSheet, Qt: 90}, {Item: *itemExcitedPhotonicMatter, Qt: 250}}, Output: []models.RecipeOutput{{Item: *itemSuperpositionOscillator, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 250}}},
		{Slug: "synthetic_power_shard", Name: "Synthetic Power Shard", Duration: 120, ProducedIn: buildingQuantumEncoder, Input: []models.RecipeInput{{Item: *itemTimeCrystal, Qt: 20}, {Item: *itemDarkMatterCrystal, Qt: 20}, {Item: *itemQuartzCrystal, Qt: 120}, {Item: *itemExcitedPhotonicMatter, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemPowerShard, Qt: 10}, {Item: *itemDarkMatterResidue, Qt: 120}}},
		{Slug: "tempered_caterium_ingot", Name: "Tempered Caterium Ingot", Duration: 80, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemCateriumOre, Qt: 60}, {Item: *itemPetroleumCoke, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCateriumIngot, Qt: 30}}},
		{Slug: "tempered_copper_ingot", Name: "Tempered Copper Ingot", Duration: 120, ProducedIn: buildingFoundry, Input: []models.RecipeInput{{Item: *itemCopperOre, Qt: 50}, {Item: *itemPetroleumCoke, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemCopperIngot, Qt: 120}}},
		{Slug: "thermal_propulsion_rocket", Name: "Thermal Propulsion Rocket", Duration: 1200, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemModularEngine, Qt: 50}, {Item: *itemTurboMotor, Qt: 20}, {Item: *itemCoolingSystem, Qt: 60}, {Item: *itemFusedModularFrame, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemThermalPropulsionRocket, Qt: 20}}},
		{Slug: "time_crystal", Name: "Time Crystal", Duration: 100, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemDiamonds, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemTimeCrystal, Qt: 10}}},
		{Slug: "turbo_blend_fuel", Name: "Turbo Blend Fuel", Duration: 80, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemFuel, Qt: 20}, {Item: *itemHeavyOilResidue, Qt: 40}, {Item: *itemSulfur, Qt: 30}, {Item: *itemPetroleumCoke, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemTurbofuel, Qt: 60}}},
		{Slug: "turbo_diamonds", Name: "Turbo Diamonds", Duration: 30, ProducedIn: buildingParicleAccelerator, Input: []models.RecipeInput{{Item: *itemCoal, Qt: 300}, {Item: *itemPackagedTurbofuel, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemDiamonds, Qt: 30}}},
		{Slug: "turbo_electric_motor", Name: "Turbo Electric Motor", Duration: 640, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemMotor, Qt: 70}, {Item: *itemRadioControlUnit, Qt: 90}, {Item: *itemElectromagneticControlRod, Qt: 50}, {Item: *itemRotor, Qt: 70}}, Output: []models.RecipeOutput{{Item: *itemTurboMotor, Qt: 30}}},
		{Slug: "turbo_heavy_fuel", Name: "Turbo Heavy Fuel", Duration: 80, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemHeavyOilResidue, Qt: 50}, {Item: *itemCompactedCoal, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemTurbofuel, Qt: 40}}},
		{Slug: "turbo_motor", Name: "Turbo Motor", Duration: 320, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemCoolingSystem, Qt: 40}, {Item: *itemRadioControlUnit, Qt: 20}, {Item: *itemMotor, Qt: 40}, {Item: *itemRubber, Qt: 240}}, Output: []models.RecipeOutput{{Item: *itemTurboMotor, Qt: 10}}},
		{Slug: "turbo_pressure_motor", Name: "Turbo Pressure Motor", Duration: 320, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemMotor, Qt: 40}, {Item: *itemPressureConversionCube, Qt: 10}, {Item: *itemPackagedNitrogenGas, Qt: 240}, {Item: *itemStator, Qt: 80}}, Output: []models.RecipeOutput{{Item: *itemTurboMotor, Qt: 20}}},
		{Slug: "turbo_rifle_ammo", Name: "Turbo Rifle Ammo", Duration: 120, ProducedIn: buildingBlender, Input: []models.RecipeInput{{Item: *itemRifleAmmo, Qt: 250}, {Item: *itemAluminumCasing, Qt: 30}, {Item: *itemTurbofuel, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemTurboRifleAmmo, Qt: 500}}},
		{Slug: "turbo_rifle_ammo", Name: "Turbo Rifle Ammo", Duration: 120, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemRifleAmmo, Qt: 250}, {Item: *itemAluminumCasing, Qt: 30}, {Item: *itemPackagedTurbofuel, Qt: 30}}, Output: []models.RecipeOutput{{Item: *itemTurboRifleAmmo, Qt: 500}}},
		{Slug: "turbofuel", Name: "Turbofuel", Duration: 160, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemFuel, Qt: 60}, {Item: *itemCompactedCoal, Qt: 40}}, Output: []models.RecipeOutput{{Item: *itemTurbofuel, Qt: 50}}},
		{Slug: "unpackage_alumina_solution", Name: "Unpackage Alumina Solution", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedAluminaSolution, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemAluminaSolution, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "unpackage_fuel", Name: "Unpackage Fuel", Duration: 20, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedFuel, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemFuel, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "unpackage_heavy_oil_residue", Name: "Unpackage Heavy Oil Residue", Duration: 60, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedHeavyOilResidue, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemHeavyOilResidue, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "unpackage_ionized_fuel", Name: "Unpackage Ionized Fuel", Duration: 30, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedIonizedFuel, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemIonizedFuel, Qt: 40}, {Item: *itemEmptyFluidTank, Qt: 20}}},
		{Slug: "unpackage_liquid_biofuel", Name: "Unpackage Liquid Biofuel", Duration: 20, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedLiquidBiofuel, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemLiquidBiofuel, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "unpackage_nitric_acid", Name: "Unpackage Nitric Acid", Duration: 30, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedNitricAcid, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemNitricAcid, Qt: 10}, {Item: *itemEmptyFluidTank, Qt: 10}}},
		{Slug: "unpackage_nitrogen_gas", Name: "Unpackage Nitrogen Gas", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedNitrogenGas, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemNitrogenGas, Qt: 40}, {Item: *itemEmptyFluidTank, Qt: 10}}},
		{Slug: "unpackage_oil", Name: "Unpackage Oil", Duration: 20, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedOil, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemCrudeOil, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "unpackage_rocket_fuel", Name: "Unpackage Rocket Fuel", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedRocketFuel, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemRocketFuel, Qt: 20}, {Item: *itemEmptyFluidTank, Qt: 10}}},
		{Slug: "unpackage_sulfuric_acid", Name: "Unpackage Sulfuric Acid", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedSulfuricAcid, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemSulfuricAcid, Qt: 10}, {Item: *itemEmptyCanister, Qt: 10}}},
		{Slug: "unpackage_turbofuel", Name: "Unpackage Turbofuel", Duration: 60, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedTurbofuel, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemTurbofuel, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "unpackage_water", Name: "Unpackage Water", Duration: 10, ProducedIn: buildingPackager, Input: []models.RecipeInput{{Item: *itemPackagedWater, Qt: 20}}, Output: []models.RecipeOutput{{Item: *itemWater, Qt: 20}, {Item: *itemEmptyCanister, Qt: 20}}},
		{Slug: "uranium_fuel_rod", Name: "Uranium Fuel Rod", Duration: 1500, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemEncasedUraniumCell, Qt: 500}, {Item: *itemEncasedIndustrialBeam, Qt: 30}, {Item: *itemElectromagneticControlRod, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemUraniumFuelRod, Qt: 10}}},
		{Slug: "uranium_fuel_rod_burning", Name: "Uranium Fuel Rod (burning)", Duration: 3000, ProducedIn: buildingNuclearGenerator, Input: []models.RecipeInput{{Item: *itemUraniumFuelRod, Qt: 10}, {Item: *itemWater, Qt: 12000}}, Output: []models.RecipeOutput{{Item: *itemUraniumWaste, Qt: 500}}},
		{Slug: "uranium_fuel_unit", Name: "Uranium Fuel Unit", Duration: 3000, ProducedIn: buildingManufacurer, Input: []models.RecipeInput{{Item: *itemEncasedUraniumCell, Qt: 1000}, {Item: *itemElectromagneticControlRod, Qt: 100}, {Item: *itemCrystalOscillator, Qt: 30}, {Item: *itemRotor, Qt: 100}}, Output: []models.RecipeOutput{{Item: *itemUraniumFuelRod, Qt: 30}}},
		{Slug: "uranium_ore_bauxite", Name: "Uranium Ore (Bauxite)", Duration: 60, ProducedIn: buildingConverter, Input: []models.RecipeInput{{Item: *itemReanimatedSAM, Qt: 10}, {Item: *itemBauxite, Qt: 480}}, Output: []models.RecipeOutput{{Item: *itemUranium, Qt: 120}}},
		{Slug: "versatile_framework", Name: "Versatile Framework", Duration: 240, ProducedIn: buildingAssember, Input: []models.RecipeInput{{Item: *itemModularFrame, Qt: 10}, {Item: *itemSteelBeam, Qt: 120}}, Output: []models.RecipeOutput{{Item: *itemVersatileFramework, Qt: 20}}},
		{Slug: "wet_concrete", Name: "Wet Concrete", Duration: 30, ProducedIn: buildingRefinery, Input: []models.RecipeInput{{Item: *itemLimestone, Qt: 60}, {Item: *itemWater, Qt: 50}}, Output: []models.RecipeOutput{{Item: *itemConcrete, Qt: 40}}},
		{Slug: "wire", Name: "Wire", Duration: 40, ProducedIn: buildingConstructor, Input: []models.RecipeInput{{Item: *itemCopperIngot, Qt: 10}}, Output: []models.RecipeOutput{{Item: *itemWire, Qt: 20}}},
	}
	err = s.DB.Save(&recipes).Error
	if err != nil {
		return err
	}

	return nil
}
