package claimDB

import "encoding/xml"

type ClaimResult struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text          string `xml:",chardata"`
		RequestMsgID  string `xml:"requestMsgID"`
		ResponseTime  string `xml:"responseTime"`
		ResponseMsgID string `xml:"responseMsgID"`
		SuccessYN     string `xml:"successYN"`
		ResultCode    string `xml:"resultCode"`
		ResultMsg     string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text string `xml:",chardata"`
		Item struct {
			Text                   string `xml:",chardata"`
			BiblioSummaryInfoArray struct {
				Text              string `xml:",chardata"`
				BiblioSummaryInfo struct {
					Text                           string `xml:",chardata"`
					ApplicationDate                string `xml:"applicationDate"`
					ApplicationNumber              string `xml:"applicationNumber"`
					ClaimCount                     string `xml:"claimCount"`
					FinalDisposal                  string `xml:"finalDisposal"`
					InventionTitle                 string `xml:"inventionTitle"`
					InventionTitleEng              string `xml:"inventionTitleEng"`
					OpenDate                       string `xml:"openDate"`
					OpenNumber                     string `xml:"openNumber"`
					OriginalApplicationDate        string `xml:"originalApplicationDate"`
					OriginalApplicationKind        string `xml:"originalApplicationKind"`
					OriginalApplicationNumber      string `xml:"originalApplicationNumber"`
					OriginalExaminationRequestDate string `xml:"originalExaminationRequestDate"`
					OriginalExaminationRequestFlag string `xml:"originalExaminationRequestFlag"`
					PublicationDate                string `xml:"publicationDate"`
					PublicationNumber              string `xml:"publicationNumber"`
					RegisterDate                   string `xml:"registerDate"`
					RegisterNumber                 string `xml:"registerNumber"`
					RegisterStatus                 string `xml:"registerStatus"`
					TranslationSubmitDate          string `xml:"translationSubmitDate"`
				} `xml:"biblioSummaryInfo"`
			} `xml:"biblioSummaryInfoArray"`
			IpcInfoArray struct {
				Text    string `xml:",chardata"`
				IpcInfo struct {
					Text      string `xml:",chardata"`
					IpcDate   string `xml:"ipcDate"`
					IpcNumber string `xml:"ipcNumber"`
				} `xml:"ipcInfo"`
			} `xml:"ipcInfoArray"`
			FamilyInfoArray struct {
				Text       string `xml:",chardata"`
				FamilyInfo string `xml:"familyInfo"`
			} `xml:"familyInfoArray"`
			AbstractInfoArray struct {
				Text         string `xml:",chardata"`
				AbstractInfo struct {
					Text      string `xml:",chardata"`
					AstrtCont string `xml:"astrtCont"`
				} `xml:"abstractInfo"`
			} `xml:"abstractInfoArray"`
			InternationalInfoArray struct {
				Text              string `xml:",chardata"`
				InternationalInfo struct {
					Text                           string `xml:",chardata"`
					InternationOpenDate            string `xml:"internationOpenDate"`
					InternationOpenNumber          string `xml:"internationOpenNumber"`
					InternationalApplicationDate   string `xml:"internationalApplicationDate"`
					InternationalApplicationNumber string `xml:"internationalApplicationNumber"`
				} `xml:"internationalInfo"`
			} `xml:"internationalInfoArray"`
			ClaimInfoArray struct {
				Text      string `xml:",chardata"`
				ClaimInfo []struct {
					Text  string `xml:",chardata"`
					Claim string `xml:"claim"`
				} `xml:"claimInfo"`
			} `xml:"claimInfoArray"`
			ApplicantInfoArray struct {
				Text          string `xml:",chardata"`
				ApplicantInfo struct {
					Text    string `xml:",chardata"`
					Address string `xml:"address"`
					Code    string `xml:"code"`
					Country string `xml:"country"`
					EngName string `xml:"engName"`
					Name    string `xml:"name"`
				} `xml:"applicantInfo"`
			} `xml:"applicantInfoArray"`
			InventorInfoArray struct {
				Text         string `xml:",chardata"`
				InventorInfo []struct {
					Text    string `xml:",chardata"`
					Address string `xml:"address"`
					Code    string `xml:"code"`
					Country string `xml:"country"`
					EngName string `xml:"engName"`
					Name    string `xml:"name"`
				} `xml:"inventorInfo"`
			} `xml:"inventorInfoArray"`
			AgentInfoArray struct {
				Text      string `xml:",chardata"`
				AgentInfo []struct {
					Text    string `xml:",chardata"`
					Address string `xml:"address"`
					Code    string `xml:"code"`
					Country string `xml:"country"`
					EngName string `xml:"engName"`
					Name    string `xml:"name"`
				} `xml:"agentInfo"`
			} `xml:"agentInfoArray"`
			PriorityInfoArray struct {
				Text         string `xml:",chardata"`
				PriorityInfo []struct {
					Text                       string `xml:",chardata"`
					PriorityApplicationCountry string `xml:"priorityApplicationCountry"`
					PriorityApplicationDate    string `xml:"priorityApplicationDate"`
					PriorityApplicationNumber  string `xml:"priorityApplicationNumber"`
				} `xml:"priorityInfo"`
			} `xml:"priorityInfoArray"`
			DesignatedStateInfoArray   string `xml:"designatedStateInfoArray"`
			PriorArtDocumentsInfoArray struct {
				Text                  string `xml:",chardata"`
				PriorArtDocumentsInfo []struct {
					Text            string `xml:",chardata"`
					DocumentsNumber string `xml:"documentsNumber"`
				} `xml:"priorArtDocumentsInfo"`
			} `xml:"priorArtDocumentsInfoArray"`
			LegalStatusInfoArray struct {
				Text            string `xml:",chardata"`
				LegalStatusInfo []struct {
					Text            string `xml:",chardata"`
					CommonCodeName  string `xml:"commonCodeName"`
					DocumentEngName string `xml:"documentEngName"`
					DocumentName    string `xml:"documentName"`
					ReceiptDate     string `xml:"receiptDate"`
					ReceiptNumber   string `xml:"receiptNumber"`
				} `xml:"legalStatusInfo"`
			} `xml:"legalStatusInfoArray"`
			ImagePathInfo struct {
				Text      string `xml:",chardata"`
				DocName   string `xml:"docName"`
				LargePath string `xml:"largePath"`
				Path      string `xml:"path"`
			} `xml:"imagePathInfo"`
		} `xml:"item"`
	} `xml:"body"`
	Count struct {
		Text       string `xml:",chardata"`
		NumOfRows  string `xml:"numOfRows"`
		PageNo     string `xml:"pageNo"`
		TotalCount string `xml:"totalCount"`
	} `xml:"count"`
}
