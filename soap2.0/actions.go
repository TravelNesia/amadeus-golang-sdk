package soap2_0

import (
	flireq051 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/air_flightinfo"
	itareq052 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/air_sellfromrecommendation"
	hsfreq073 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/command_cryptic"
	ttktiq091 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/docissuance_issueticket"
	farqnq071 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_checkrules"
	fcuqcq081 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_convertcurrency"
	tibnrq124 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_informativebestpricingwithoutpnr"
	tipnrq124 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_informativepricingwithoutpnr"
	fmpcaq143 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricercalendar"
	fmpcaq122 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricercalendar_old"
	fmptbq143 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch"
	fmptbq163 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch_new"
	fmptbq123 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch_old"
	tpcbrq141 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_pricepnrwithbookingclass"
	tpcbrq124 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_pricepnrwithbookingclass_old"
	pnradd113 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/pnr_addmultielements"
	pnrxcl113 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/pnr_cancel"
	pnrret113 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/pnr_retrieve"
	qcddrq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_countplanner"
	qcsdrq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_counttotal"
	qdqlrq111 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_list"
	quqmuq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_moveitem"
	quqpcq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_placepnr"
	quqmdq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_removeitem"
	vlsslq061 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/security_authenticate"
	vlssoq041 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/security_signout"
	tautcq041 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_createtstfrompricing"
	ccvrqt061 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_creditcardcheck"
	ttstdq041 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_deletetst"
	ttstrq071 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_displaytst"

	flires051 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/air_flightinfo_reply"
	itares052 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/air_sellfromrecommendation_reply"
	hsfres073 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/command_cryptic_reply"
	ttktir091 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/docissuance_issueticket_reply"
	farqnr071 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_checkrules_reply"
	fcuqcr081 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_convertcurrency_reply"
	tibnrr124 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_informativebestpricingwithoutpnr_reply"
	tipnrr124 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_informativepricingwithoutpnr_reply"
	fmpcar122 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricercalendar_old_reply"
	fmpcar143 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricercalendar_reply"
	fmptbr163 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_new_reply"
	fmptbr123 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_old_reply"
	fmptbr143 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_reply"
	tpcbrr124 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_pricepnrwithbookingclass_old_reply"
	tpcbrr141 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_pricepnrwithbookingclass_reply"
	tnlres001 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/pnr_list"
	pnracc113 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/pnr_reply"
	qcddrr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_countplanner_reply"
	qcsdrr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_counttotal_reply"
	qdqlrr111 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_list_reply"
	quqmur031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_moveitem_reply"
	quqpcr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_placepnr_reply"
	quqmdr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_removeitem_reply"
	vlsslr061 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/security_authenticate_reply"
	vlssor041 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/security_signout_reply"
	tautcr041 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_createtstfrompricing_reply"
	ccvrsp061 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_creditcardcheck_reply"
	ttstdr041 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_deletetst_reply"
	ttstrr071 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_displaytst_reply"
)

var url = "http://webservices.amadeus.com"

func (service *WebServicesPT) AirFlightInfo(query *flireq051.AirFlightInfo) (*flires051.AirFlightInfoReply, error) {
	soapAction := "FLIREQ_05_1_1A"

	reply := new(flires051.AirFlightInfoReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) AirSellFromRecommendation(query *itareq052.AirSellFromRecommendation) (*itares052.AirSellFromRecommendationReply, error) {
	soapAction := "ITAREQ_05_2_IA"

	reply := new(itares052.AirSellFromRecommendationReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) CommandCryptic(query *hsfreq073.CommandCryptic) (*hsfres073.CommandCrypticReply, *Session, error) {
	soapAction := "HSFREQ_07_3_1A"

	reply := new(hsfres073.CommandCrypticReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, nil, err
	}

	return reply, session, nil
}

func (service *WebServicesPT) DocIssuanceIssueTicket(query *ttktiq091.DocIssuanceIssueTicket) (*ttktir091.DocIssuanceIssueTicketReply, error) {
	soapAction := "TTKTIQ_09_1_1A"

	reply := new(ttktir091.DocIssuanceIssueTicketReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareCheckRules(query *farqnq071.FareCheckRules) (*farqnr071.FareCheckRulesReply, error) {
	soapAction := "FARQNQ_07_1_1A"

	reply := new(farqnr071.FareCheckRulesReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareConvertCurrency(query *fcuqcq081.FareConvertCurrency) (*fcuqcr081.FareConvertCurrencyReply, error) {
	soapAction := "FCUQCQ_08_1_1A"

	reply := new(fcuqcr081.FareConvertCurrencyReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareInformativeBestPricingWithoutPNR(query *tibnrq124.FareInformativeBestPricingWithoutPNR) (*tibnrr124.FareInformativeBestPricingWithoutPNRReply, error) {
	soapAction := "TIBNRQ_12_4_1A"

	reply := new(tibnrr124.FareInformativeBestPricingWithoutPNRReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareInformativePricingWithoutPNR(query *tipnrq124.FareInformativePricingWithoutPNR) (*tipnrr124.FareInformativePricingWithoutPNRReply, error) {
	soapAction := "TIPNRQ_12_4_1A"

	reply := new(tipnrr124.FareInformativePricingWithoutPNRReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareMasterPricerCalendar(query *fmpcaq143.FareMasterPricerCalendar) (*fmpcar143.FareMasterPricerCalendarReply, error) {
	soapAction := "FMPCAQ_14_3_1A"

	reply := new(fmpcar143.FareMasterPricerCalendarReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareMasterPricerCalendarOld(query *fmpcaq122.FareMasterPricerCalendarOld) (*fmpcar122.FareMasterPricerCalendarOldReply, error) {
	soapAction := "FMPCAQ_12_2_1A"

	reply := new(fmpcar122.FareMasterPricerCalendarOldReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareMasterPricerTravelBoardSearchNew(query *fmptbq163.FareMasterPricerTravelBoardSearch) (*fmptbr163.FareMasterPricerTravelBoardSearchReply, error) {
	soapAction := "FMPTBQ_16_3_1A"

	reply := new(fmptbr163.FareMasterPricerTravelBoardSearchReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareMasterPricerTravelBoardSearch(query *fmptbq143.FareMasterPricerTravelBoardSearch) (*fmptbr143.FareMasterPricerTravelBoardSearchReply, error) {
	soapAction := "FMPTBQ_14_3_1A"

	reply := new(fmptbr143.FareMasterPricerTravelBoardSearchReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FareMasterPricerTravelBoardSearchOld(query *fmptbq123.FareMasterPricerTravelBoardSearchOld) (*fmptbr123.FareMasterPricerTravelBoardSearchOldReply, error) {
	soapAction := "FMPTBQ_12_3_1A"

	reply := new(fmptbr123.FareMasterPricerTravelBoardSearchOldReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FarePricePNRWithBookingClass(query *tpcbrq141.FarePricePNRWithBookingClass) (*tpcbrr141.FarePricePNRWithBookingClassReply, error) {
	soapAction := "TPCBRQ_14_1_1A"

	reply := new(tpcbrr141.FarePricePNRWithBookingClassReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) FarePricePNRWithBookingClassOld(query *tpcbrq124.FarePricePNRWithBookingClassOld) (*tpcbrr124.FarePricePNRWithBookingClassOldReply, error) {
	soapAction := "TPCBRQ_12_4_1A"

	reply := new(tpcbrr124.FarePricePNRWithBookingClassOldReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) PNRAddMultiElements(query *pnradd113.PNRAddMultiElements) (*pnracc113.PNRReply, error) {
	soapAction := "PNRADD_11_3_1A"

	reply := new(pnracc113.PNRReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) PNRCancel(query *pnrxcl113.PNRCancel) (*pnracc113.PNRReply, error) {
	soapAction := "PNRXCL_11_3_1A"

	reply := new(pnracc113.PNRReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) PNRRetrieve(query *pnrret113.PNRRetrieve) (*pnracc113.PNRReply, error) {
	soapAction := "PNRRET_11_3_1A"

	reply := new(pnracc113.PNRReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) PNRRetrieve2(query *pnrret113.PNRRetrieve) (*tnlres001.PNRList, error) {
	soapAction := "PNRRET_11_3_1A"

	reply := new(tnlres001.PNRList)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) QueueCountPlanner(query *qcddrq031.QueueCountPlanner) (*qcddrr031.QueueCountPlannerReply, error) {
	soapAction := "QCDDRQ_03_1_1A"

	reply := new(qcddrr031.QueueCountPlannerReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) QueueCountTotal(query *qcsdrq031.QueueCountTotal) (*qcsdrr031.QueueCountTotalReply, error) {
	soapAction := "QCSDRQ_03_1_1A"

	reply := new(qcsdrr031.QueueCountTotalReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) QueueList(query *qdqlrq111.QueueList) (*qdqlrr111.QueueListReply, error) {
	soapAction := "QDQLRQ_11_1_1A"

	reply := new(qdqlrr111.QueueListReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) QueueMoveItem(query *quqmuq031.QueueMoveItem) (*quqmur031.QueueMoveItemReply, error) {
	soapAction := "QUQMUQ_03_1_1A"

	reply := new(quqmur031.QueueMoveItemReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) QueuePlacePNR(query *quqpcq031.QueuePlacePNR) (*quqpcr031.QueuePlacePNRReply, error) {
	soapAction := "QUQPCQ_03_1_1A"

	reply := new(quqpcr031.QueuePlacePNRReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) QueueRemoveItem(query *quqmdq031.QueueRemoveItem) (*quqmdr031.QueueRemoveItemReply, error) {
	soapAction := "QUQMDQ_03_1_1A"

	reply := new(quqmdr031.QueueRemoveItemReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) SecurityAuthenticate(query *vlsslq061.SecurityAuthenticate) (*vlsslr061.SecurityAuthenticateReply, *Session, error) {
	soapAction := "VLSSLQ_06_1_1A"

	reply := new(vlsslr061.SecurityAuthenticateReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, nil, err
	}

	return reply, session, nil
}

func (service *WebServicesPT) SecuritySignOut(query *vlssoq041.SecuritySignOut) (*vlssor041.SecuritySignOutReply, error) {
	soapAction := "VLSSOQ_04_1_1A"

	reply := new(vlssor041.SecuritySignOutReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) TicketCreateTSTFromPricing(query *tautcq041.TicketCreateTSTFromPricing) (*tautcr041.TicketCreateTSTFromPricingReply, error) {
	soapAction := "TAUTCQ_04_1_1A"

	reply := new(tautcr041.TicketCreateTSTFromPricingReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) TicketCreditCardCheck(query *ccvrqt061.TicketCreditCardCheck) (*ccvrsp061.TicketCreditCardCheckReply, error) {
	soapAction := "CCVRQT_06_1_1A"

	reply := new(ccvrsp061.TicketCreditCardCheckReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) TicketDeleteTST(query *ttstdq041.TicketDeleteTST) (*ttstdr041.TicketDeleteTSTReply, error) {
	soapAction := "TTSTDQ_04_1_1A"

	reply := new(ttstdr041.TicketDeleteTSTReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (service *WebServicesPT) TicketDisplayTST(query *ttstrq071.TicketDisplayTST) (*ttstrr071.TicketDisplayTSTReply, error) {
	soapAction := "TTSTRQ_07_1_1A"

	reply := new(ttstrr071.TicketDisplayTSTReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, query, reply, session)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
