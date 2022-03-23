package models

// Ncpdp
// I got this fixed-width definition from Medsavvy's existing NCPDP importer.
// Not sure where (or even if) this is documented in a publicly available place.
// TODO: Check some of the types. I just made them all strings for now because of leading zeros in IDs
type Ncpdp struct {
    NcpdpProviderID           string `fixed:"1,7" dynamodbav:"ncpdp_provider_id"`
    LegalBusinessName         string `fixed:"8,67" dynamodbav:"legal_business_name"`
    ProviderName              string `fixed:"68,127" dynamodbav:"provider_name"`
    DoctorName                string `fixed:"128,187" dynamodbav:"doctor_name"`
    StoreNumber               string `fixed:"188,197" dynamodbav:"store_number"`
    StoreAddress1             string `fixed:"198,252" dynamodbav:"store_address_1"`
    StoreAddress2             string `fixed:"253,307" dynamodbav:"store_address_2"`
    StoreCity                 string `fixed:"308,337" dynamodbav:"store_city"`
    StoreStateCode            string `fixed:"338,339" dynamodbav:"store_state_code"`
    StoreZipCode              string `fixed:"340,348" dynamodbav:"store_zip_code"`
    StorePhoneNumber          string `fixed:"349,358" dynamodbav:"store_phone_number"`
    StoreExtension            string `fixed:"359,363" dynamodbav:"store_extension"`
    StoreFax                  string `fixed:"364,373" dynamodbav:"store_fax"`
    StoreEmailAddress         string `fixed:"374,423" dynamodbav:"store_email_address"`
    StoreDirections           string `fixed:"424,473" dynamodbav:"store_directions"`
    StoreCounty               string `fixed:"474,478" dynamodbav:"store_county"`
    StoreMsa                  string `fixed:"479,482" dynamodbav:"store_msa"`
    StorePmsa                 string `fixed:"482,486" dynamodbav:"store_pmsa"`
    Store24hrOperationFlag    string `fixed:"487,487" dynamodbav:"store_24_hr_operation_flag"`
    StoreProviderHours        string `fixed:"488,522" dynamodbav:"store_provider_hours"`
    StoreVotingDistrict       string `fixed:"523,526" dynamodbav:"store_voting_district"`
    StoreLanguageCode1        string `fixed:"527,528" dynamodbav:"store_language_code_1"`
    StoreLanguageCode2        string `fixed:"529,530" dynamodbav:"store_language_code_2"`
    StoreLanguageCode3        string `fixed:"531,532" dynamodbav:"store_language_code_3"`
    StoreLanguageCode4        string `fixed:"533,534" dynamodbav:"store_language_code_4"`
    StoreLanguageCode5        string `fixed:"535,536" dynamodbav:"store_language_code_5"`
    StoreOpenDate             string `fixed:"537,544" dynamodbav:"store_open_date"`
    StoreClosureDate          string `fixed:"545,552" json:"store_closure_date"`
    MailingAddress1           string `fixed:"553,607" json:"mailing_address_1"`
    MailingAddress2           string `fixed:"608,662" json:"mailing_address_2"`
    MailingCity               string `fixed:"663,692" json:"mailing_city"`
    MailingStateCode          string `fixed:"693,694" json:"mailing_state_code"`
    MailingZipCode            string `fixed:"695,703" json:"mailing_zip_code"`
    ContactLastName           string `fixed:"704,723" json:"contact_last_name"`
    ContactFirstName          string `fixed:"724,743" json:"contact_first_name"`
    ContactMiddleInitial      string `fixed:"744,744" json:"contact_middle_initial"`
    ContactTitle              string `fixed:"745,774" json:"contact_title"`
    ContactPhoneNumber        string `fixed:"775,784" json:"contact_phone_number"`
    ContactExtension          string `fixed:"785,789" json:"contact_extension"`
    ContactEmailAddress       string `fixed:"790,839" json:"contact_email_address"`
    DispenserClassCode        string `fixed:"840,841" json:"dispenser_class_code"`
    PrimaryProviderTypeCode   string `fixed:"842,843" json:"primary_provider_type_code"`
    SecondaryProviderTypeCode string `fixed:"844,845" json:"secondary_provider_type_code"`
    TertiaryProviderTypeCode  string `fixed:"846,847" json:"tertiary_provider_type_code"`
    MedicareProviderID        string `fixed:"848,857" json:"medicare_provider_id"`
    NationalProviderID        string `fixed:"858,867" json:"national_provider_id"`
    DeaRegistrationID         string `fixed:"868,879" json:"dea_registration_id"`
    DeaExpirationDate         string `fixed:"880,887" json:"dea_expiration_date"`
    FederalTaxIDNumber        string `fixed:"888,902" json:"federal_tax_id_number"`
    StateIncomeTaxIDNumber    string `fixed:"903,917" json:"state_income_tax_id_number"`
    DeactivationCode          string `fixed:"918,919" json:"deactivation_code"`
    ReinstatementCode         string `fixed:"920,921" json:"reinstatement_code"`
    ReinstatementDate         string `fixed:"922,929" json:"reinstatement_date"`
    TransactionCode           string `fixed:"930,930" json:"transaction_code"`
    TransactionDate           string `fixed:"931,938" json:"transaction_date"`
}
