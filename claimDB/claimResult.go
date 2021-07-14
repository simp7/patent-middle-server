package claimDB

import "encoding/xml"

type ClaimResult struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text  string `xml:",chardata"`
		Items struct {
			Text              string `xml:",chardata"`
			AbstractInfoArray struct {
				Text         string `xml:",chardata"`
				Type         string `xml:"type,attr"`
				AbstractInfo struct {
					Text      string `xml:",chardata"`
					Type      string `xml:"type,attr"`
					AstrtCont string `xml:"astrtCont"`
				} `xml:"abstractInfo"`
			} `xml:"abstractInfoArray"`
			AgentInfoArray struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				AgentInfo struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Address string `xml:"address"`
					Code    string `xml:"code"`
					Country string `xml:"country"`
					EngName string `xml:"engName"`
					Name    string `xml:"name"`
				} `xml:"agentInfo"`
			} `xml:"agentInfoArray"`
			ApplicantInfoArray struct {
				Text          string `xml:",chardata"`
				Type          string `xml:"type,attr"`
				ApplicantInfo struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Address string `xml:"address"`
					Code    string `xml:"code"`
					Country string `xml:"country"`
					EngName string `xml:"engName"`
					Name    string `xml:"name"`
				} `xml:"applicantInfo"`
			} `xml:"applicantInfoArray"`
			BiblioSummaryInfoArray struct {
				Text              string `xml:",chardata"`
				Type              string `xml:"type,attr"`
				BiblioSummaryInfo struct {
					Text                           string `xml:",chardata"`
					Type                           string `xml:"type,attr"`
					ApplicationDate                string `xml:"applicationDate"`
					ApplicationFlag                string `xml:"applicationFlag"`
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
			ClaimInfoArray struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				ClaimInfo []struct {
					Text  string `xml:",chardata"`
					Type  string `xml:"type,attr"`
					Claim string `xml:"claim"`
				} `xml:"claimInfo"`
			} `xml:"claimInfoArray"`
			DesignatedStateInfoArray struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"designatedStateInfoArray"`
			FamilyInfoArray struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				FamilyInfo struct {
					Text                    string `xml:",chardata"`
					Type                    string `xml:"type,attr"`
					FamilyApplicationNumber struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
						Xsi  string `xml:"xsi,attr"`
					} `xml:"familyApplicationNumber"`
				} `xml:"familyInfo"`
			} `xml:"familyInfoArray"`
			ImagePathInfo struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				DocName   string `xml:"docName"`
				LargePath string `xml:"largePath"`
				Path      string `xml:"path"`
			} `xml:"imagePathInfo"`
			InternationalInfoArray struct {
				Text              string `xml:",chardata"`
				Type              string `xml:"type,attr"`
				InternationalInfo struct {
					Text                           string `xml:",chardata"`
					Type                           string `xml:"type,attr"`
					InternationOpenDate            string `xml:"internationOpenDate"`
					InternationOpenNumber          string `xml:"internationOpenNumber"`
					InternationalApplicationDate   string `xml:"internationalApplicationDate"`
					InternationalApplicationNumber struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
						Xsi  string `xml:"xsi,attr"`
					} `xml:"internationalApplicationNumber"`
				} `xml:"internationalInfo"`
			} `xml:"internationalInfoArray"`
			InventorInfoArray struct {
				Text         string `xml:",chardata"`
				Type         string `xml:"type,attr"`
				InventorInfo []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Address string `xml:"address"`
					Code    string `xml:"code"`
					Country string `xml:"country"`
					EngName string `xml:"engName"`
					Name    string `xml:"name"`
				} `xml:"inventorInfo"`
			} `xml:"inventorInfoArray"`
			IpcInfoArray struct {
				Text    string `xml:",chardata"`
				Type    string `xml:"type,attr"`
				IpcInfo []struct {
					Text      string `xml:",chardata"`
					Type      string `xml:"type,attr"`
					IpcDate   string `xml:"ipcDate"`
					IpcNumber string `xml:"ipcNumber"`
				} `xml:"ipcInfo"`
			} `xml:"ipcInfoArray"`
			LegalStatusInfoArray struct {
				Text            string `xml:",chardata"`
				Type            string `xml:"type,attr"`
				LegalStatusInfo []struct {
					Text            string `xml:",chardata"`
					Type            string `xml:"type,attr"`
					CommonCodeName  string `xml:"commonCodeName"`
					DocumentEngName string `xml:"documentEngName"`
					DocumentName    string `xml:"documentName"`
					ReceiptDate     string `xml:"receiptDate"`
					ReceiptNumber   string `xml:"receiptNumber"`
				} `xml:"legalStatusInfo"`
			} `xml:"legalStatusInfoArray"`
			PriorArtDocumentsInfoArray struct {
				Text                  string `xml:",chardata"`
				Type                  string `xml:"type,attr"`
				PriorArtDocumentsInfo []struct {
					Text            string `xml:",chardata"`
					Type            string `xml:"type,attr"`
					DocumentsNumber string `xml:"documentsNumber"`
				} `xml:"priorArtDocumentsInfo"`
			} `xml:"priorArtDocumentsInfoArray"`
			PriorityInfoArray struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"priorityInfoArray"`
		} `xml:"items"`
	} `xml:"body"`
}
