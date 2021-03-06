package models

type Item struct {
	FSID                             string   `json:"fs_id"`
	URL                              string   `json:"url"`
	Type                             string   `json:"type"`
	ClubNintendo                     bool     `json:"club_nintendo"`
	PrettyDateS                      string   `json:"pretty_date_s"`
	PlayModeTvModeB                  bool     `json:"play_mode_tv_mode_b"`
	PlayModeHandheldModeB            bool     `json:"play_mode_handheld_mode_b"`
	DemoAvailability                 *bool    `json:"demo_availability,omitempty"`
	CompatibleController             []string `json:"compatible_controller"`
	ImageURL                         string   `json:"image_url"`
	PaidSubscriptionRequiredB        bool     `json:"paid_subscription_required_b"`
	CloudSavesB                      bool     `json:"cloud_saves_b"`
	DigitalVersionB                  bool     `json:"digital_version_b"`
	TitleExtrasTxt                   []string `json:"title_extras_txt"`
	ImageURLH2X1S                    string   `json:"image_url_h2x1_s"`
	GameCategoriesTxt                []string `json:"game_categories_txt"`
	PlayModeTabletopModeB            bool     `json:"play_mode_tabletop_mode_b"`
	Publisher                        string   `json:"publisher"`
	Excerpt                          string   `json:"excerpt"`
	LanguageAvailability             []string `json:"language_availability"`
	PriceHasDiscountB                bool     `json:"price_has_discount_b"`
	PriceDiscountPercentageF         float64  `json:"price_discount_percentage_f"`
	Title                            string   `json:"title"`
	PlayersTo                        int64    `json:"players_to"`
	WishlistEmailBanner640WImageURLS string   `json:"wishlist_email_banner640w_image_url_s"`
	PrettyGameCategoriesTxt          []string `json:"pretty_game_categories_txt"`
	PrettyAgeratingS                 string   `json:"pretty_agerating_s"`
	PriceRegularF                    float64  `json:"price_regular_f"`
	EshopRemovedB                    bool     `json:"eshop_removed_b"`
	PriceLowestF                     float64  `json:"price_lowest_f"`
	PhysicalVersionB                 bool     `json:"physical_version_b"`
	Version                          float64  `json:"_version_"`
	//ChangeDate                       string   `json:"change_date"`
	//DatesReleasedDts                 []string `json:"dates_released_dts"`
	//GameSeriesTxt                    []string `json:"game_series_txt"`
	//ProductCodeTxt                   []string `json:"product_code_txt"`
	//ImageURLSqS                      string   `json:"image_url_sq_s"`
	//DeprioritiseB                    bool     `json:"deprioritise_b"`
	//GameSeriesT                      *string  `json:"game_series_t,omitempty"`
	//PGS                              string   `json:"pg_s"`
	//GiftFinderDetailPageImageURLS    string   `json:"gift_finder_detail_page_image_url_s"`
	//OriginallyForT                   string   `json:"originally_for_t"`
	//Priority                         string   `json:"priority"`
	//SystemType                       []string `json:"system_type"`
	//AgeRatingSortingI                int64    `json:"age_rating_sorting_i"`
	//ProductCodeSs                    []string `json:"product_code_ss"`
	//NsuidTxt                         []string `json:"nsuid_txt"`
	//DateFrom                         string   `json:"date_from"`
	//SortingTitle                     string   `json:"sorting_title"`
	//GiftFinderCarouselImageURLS      string   `json:"gift_finder_carousel_image_url_s"`
	//WishlistEmailSquareImageURLS     string   `json:"wishlist_email_square_image_url_s"`
	//PlayableOnTxt                    []string `json:"playable_on_txt"`
	//HitsI                            int64    `json:"hits_i"`
	//GiftFinderWishlistImageURLS      string   `json:"gift_finder_wishlist_image_url_s"`
	//SwitchGameVoucherB               bool     `json:"switch_game_voucher_b"`
	//GameCategory                     []string `json:"game_category"`
	//SystemNamesTxt                   []string `json:"system_names_txt"`
	//AgeRatingType                    string   `json:"age_rating_type"`
	//GiftFinderDetailPageStoreLinkS   string   `json:"gift_finder_detail_page_store_link_s"`
	//PriceSortingF                    float64  `json:"price_sorting_f"`
	//AgeRatingValue                   string   `json:"age_rating_value"`
	//Developer                        *string  `json:"developer,omitempty"`
	//WishlistEmailBanner460WImageURLS string   `json:"wishlist_email_banner460w_image_url_s"`
	//Popularity                       int64    `json:"popularity"`
	//VoiceChatB                       *bool    `json:"voice_chat_b,omitempty"`
	//PlayersFrom                      *int64   `json:"players_from,omitempty"`
}
