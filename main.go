package main

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct{
	 Id uuid.UUID
	 UserRoles string
	 ContactInfo ContactInfo
		
}
type ContactInfo struct{
	Id uuid.UUID
	FirstName string
	LastName string
	Phone string
	Email string
	StreetAddress string
	UnitNumber int
	City string
	State string
	Zip int
	LastLoginDate time.Time
	LastLoginIpAddress time.Time
	CreationDate time.Time
}
// type UserRole struct{
// 	SuperAdmin
// 	Admin
// 	Reseller
// 	Subscriber
// }
type Domain struct{
	 Id uuid.UUID
	 DomainUrl string
	 DomainType string
		 // Main
		 // Alias
		 // Parked
}
type BusinessInfo struct{
	 Id uuid.UUID
	 BusinessName string
	 BusinessLogo string
	 BusinessSlogan string
	 BusinessTypes []string
	PaymentTypes []string
		// Cash
		// Check
		// MoneyOrder
		// CreditCard
		// Visa
		// MasterCard
		// AmericanExpress
}
type Location struct{
	 ID uuid.UUID
	 ContactInfo ContactInfo
	 Street string
	 Unit string
	 City string
	 State string
}
type SocialMedia struct{
	 YoutubeUrl  string
	 YoutubeIcon  string
	 FacebookUrl  string
	 FacebookIcon  string
	 LinkedInUrl  string
	 LinkedInIcon  string
	 TwitterUrl  string
	 TwitterIcon  string
	 PinterestUrl  string
	 PinterestIcon  string
	 SnapChatUrl  string
	 SnapChatIcon  string
	 GooglePlusUrl  string
	 GooglePlusIcon  string
}
type BusinessDirectories struct{
	 BbbUrl  string
	 YelpUrl  string
	 AngiesListUrl  string
	 HomeAdvisorUrl  string
}
type Tools struct{
	 GoogleMyBusiness  string
	 GoogleAnalytics  string
	 GoogleTagManager  string
	 GoogleSearchConsole  string
	 GoogleAdwords  string
	 BingWebmasters  string
	 KeywordExtractor string
}
type Service struct{
	 Id uuid.UUID
	 Name string
	 Title string
	 ShortDescription string
	 LongDescription string
	 FeaturedImageUrl string
	 MainHeading string
	 SubHeading string
	 LinkText string
	 LinkUrl string
	 Benefits []string
	 Licensed bool
	 Bonded bool
	 Insured bool
	 Certifications []string
	 YearsExperience time.Time
	 SimilarService []Service
}
type Awards struct{
	 Id uuid.UUID
	 Title string
	 Image string
	 Location Location
	 Date time.Time
}
type Reviews struct{
	 Id uuid.UUID
	 State string
	 City string
	 Latitude float64
	 Longitude float64
	 AuthorName string
	 AuthorEmail string
	 AuthorPhone string
	 Summary string
	 Detail string
	 Date time.Time
	 QualityRating int
	 PriceRating int
	 ConvenienceRating int
}
type Checkins struct{
	 Id uuid.UUID
	 UserId uuid.UUID
	 DisplayEmail string
	 DisplayName string
	 CheckinDateTime time.Time
	 ServicePerformed string
	 Street string
	 City string
	 State string
	 Zip int
	 Latitude float64
	 Longitude float64
	 ImageUrl string
	 SentReviewRequest bool
	 ReviewCompleted bool
	 ReviewId uuid.UUID
	 CustomerName string
	 CustomerPhone string
}
type AreasServed struct{
	 Id uuid.UUID
	 ZipCodeCenter int
	 MileRadius int
	 arealist []string
	 zip int
	 state string
	 county string
	 city string
	 latitude float64
	 longitude float64
}
type Deals struct{
	 Id uuid.UUID
	 Offer string
	 Duration time.Time
	 StartDate time.Time
	 EndDate time.Time
	 Requirements []string
}
type Tags struct{
	 Id uuid.UUID
	 Name string
	 ShortDescription string
	 LongDescription string
	 Thumbnail string
	 ThumbnailAltText string
}
type Albums struct{
	 Id uuid.UUID
	 Name string
	 Descriptiong string
	 Thumbnail string
	
}
type Galleries struct{
	
	 Id uuid.UUID
	 Name string
	 Description string
	 Thumbnail string
}
type Images struct{
	 Id uuid.UUID
	 Url string
	 Caption string
	 AltText string
}
type Guarantees struct{
	 Id uuid.UUID
	 Title string
	 ShortDescription string
	 LongDescription string
	 Benefits []string
}
type Warranties struct{
	 Id uuid.UUID
	 Duration time.Month
	 Coverage Service
	 ContactEmail string
}
func main(){

	routes := []string{
		"/api/:token/:account/",
		"/api/:token/domains/:domainid/",
		"/api/:token/generalbusinessinfo/:businessid/",
		"/api/:token/locations/:locationid",
		"/api/:token/social/",
		"/api/:token/onlinebusinessdirectories/",
		"/api/:token/webtools/",
		"/api/:token/services/:serviceid",
		"/api/:token/awards/:awardid",
		"/api/:token/reviews/:reviewid",
		"/api/:token/checkins/:checkinid",
		"/api/:token/areasserved/",
		"/api/:token/deals/",
		"/api/:token/tags/:tagid",
		"/api/:token/albums/:albumid",
		"/api/:token/galleries/:galleryid",
		"/api/:token/images/:imageid",
		"/api/:token/guarantees/:guaranteeid",
		"/api/:token/warranties/:warrantyid",
	}
	
	for _,r := range routes{
		println(r)
	}
}



