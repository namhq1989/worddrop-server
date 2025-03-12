package scraper

const (
	// Politics
	politicoDomainName    = "www.politico.com"
	theHillDomainName     = "thehill.com"
	fiveThirtyEightDomain = "fivethirtyeight.com"
	realClearPolitics     = "www.realclearpolitics.com"
	axiosDomain           = "www.axios.com"
	voxPoliticsDomain     = "www.vox.com"

	// Technology
	techCrunchDomain    = "techcrunch.com"
	wiredDomain         = "www.wired.com"
	arsTechnicaDomain   = "arstechnica.com"
	thevergeDomain      = "www.theverge.com"
	mitTechReviewDomain = "www.technologyreview.com"
	cnetDomain          = "www.cnet.com"

	// Business
	bloombergDomain       = "www.bloomberg.com"
	forbesDomain          = "www.forbes.com"
	cnbcDomain            = "www.cnbc.com"
	wsjDomain             = "www.wsj.com"
	financialTimesDomain  = "www.ft.com"
	businessInsiderDomain = "www.businessinsider.com"

	// Science
	scientificAmericanDomain = "www.scientificamerican.com"
	scienceDailyDomain       = "www.sciencedaily.com"
	natureDomain             = "www.nature.com"
	scienceOrgDomain         = "www.science.org"
	quantaMagazineDomain     = "www.quantamagazine.org"
	liveScienceDomain        = "www.livescience.com"

	// Health
	webmdDomain            = "www.webmd.com"
	harvardHealthDomain    = "www.health.harvard.edu"
	medicalNewsTodayDomain = "www.medicalnewstoday.com"
	healthlineDomain       = "www.healthline.com"
	cdcDomain              = "www.cdc.gov"
	mayoClinicDomain       = "www.mayoclinic.org"

	// Sports
	espnDomain              = "www.espn.com"
	sportsIllustratedDomain = "www.si.com"
	theAthleticDomain       = "theathletic.com"
	bleacherReportDomain    = "bleacherreport.com"
	yahooSportsDomain       = "sports.yahoo.com"
	cbsSportsDomain         = "www.cbssports.com"

	// Entertainment
	varietyDomain              = "variety.com"
	theHollywoodReporterDomain = "www.hollywoodreporter.com"
	entertainmentWeeklyDomain  = "ew.com"
	deadlineDomain             = "deadline.com"
	ignDomain                  = "www.ign.com"
	vultureDomain              = "www.vulture.com"

	// World News
	bbcNewsDomain       = "www.bbc.com"
	alJazeeraDomain     = "www.aljazeera.com"
	reutersDomain       = "www.reuters.com"
	theGuardianDomain   = "www.theguardian.com"
	france24Domain      = "www.france24.com"
	deutscheWelleDomain = "www.dw.com"

	// Education
	chronicleHigherEdDomain = "www.chronicle.com"
	insideHigherEdDomain    = "www.insidehighered.com"
	edSurgeDomain           = "www.edsurge.com"
	educationWeekDomain     = "www.edweek.org"
	hechingerReportDomain   = "hechingerreport.org"
	timesHigherEdDomain     = "www.timeshighereducation.com"
)

func websitesByCategory() map[CategoryType][]string {
	return map[CategoryType][]string{
		CategoryPolitics: {
			politicoDomainName,
			theHillDomainName,
			fiveThirtyEightDomain,
			realClearPolitics,
			axiosDomain,
			voxPoliticsDomain,
		},
		CategoryTechnology: {
			techCrunchDomain,
			wiredDomain,
			arsTechnicaDomain,
			thevergeDomain,
			mitTechReviewDomain,
			cnetDomain,
		},
		CategoryBusiness: {
			bloombergDomain,
			forbesDomain,
			cnbcDomain,
			wsjDomain,
			financialTimesDomain,
			businessInsiderDomain,
		},
		CategoryScience: {
			scientificAmericanDomain,
			scienceDailyDomain,
			natureDomain,
			scienceOrgDomain,
			quantaMagazineDomain,
			liveScienceDomain,
		},
		CategoryHealth: {
			webmdDomain,
			harvardHealthDomain,
			medicalNewsTodayDomain,
			healthlineDomain,
			cdcDomain,
			mayoClinicDomain,
		},
		CategorySports: {
			espnDomain,
			sportsIllustratedDomain,
			theAthleticDomain,
			bleacherReportDomain,
			yahooSportsDomain,
			cbsSportsDomain,
		},
		CategoryEntertainment: {
			varietyDomain,
			theHollywoodReporterDomain,
			entertainmentWeeklyDomain,
			deadlineDomain,
			ignDomain,
			vultureDomain,
		},
		CategoryWorldNews: {
			bbcNewsDomain,
			alJazeeraDomain,
			reutersDomain,
			theGuardianDomain,
			france24Domain,
			deutscheWelleDomain,
		},
		CategoryEducation: {
			chronicleHigherEdDomain,
			insideHigherEdDomain,
			edSurgeDomain,
			educationWeekDomain,
			hechingerReportDomain,
			timesHigherEdDomain,
		},
	}
}
