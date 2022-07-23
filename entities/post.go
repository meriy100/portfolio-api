package entities

import "strings"

type Post struct {
	BodyMd string `json:"body_md"`
}

func (p *Post) ToProfile() (*Profile, error) {
	parts := strings.Split(p.BodyMd, "##")

	var job string
	var description string

	for _, part := range parts {
		if strings.HasPrefix(part, " job") {
			job = strings.Replace(part, "job\r\n", "", 1)
		}

		if strings.HasPrefix(part, " description") {
			description = strings.Replace(part, "description\r\n", "", 1)
		}

	}

	return NewProfile(job, description), nil
}

//func FetchPost() ([]byte, error) {
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", "https://api.esa.io/v1/teams/meriy100/posts/253", nil)
//	if err != nil {
//		return []byte{}, err
//	}
//	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ESA_ACCESS_TOKEN")))
//	resp, err := client.Do(req)
//	if err != nil {
//		return []byte{}, err
//	}
//	defer resp.Body.Close()
//
//	byteArray, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return []byte{}, err
//	}
//	return byteArray, nil
//}

//func getFireBaseClient(ctx context.Context) (*firestore.Client, error) {
//	opt := option.WithCredentialsFile("serverless_function_source_code/serviceAccountKey.json")
//	cur, _ := os.Getwd()
//	fmt.Println(cur)
//
//	fileInfos, err := ioutil.ReadDir(cur)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, fileInfo := range fileInfos {
//		fmt.Println(fileInfo.Name())
//	}
//
//	fmt.Printf("opt : %v\n", opt)
//	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "portfolio-357112"}, opt)
//	if err != nil {
//		fmt.Printf("Failed firebase.NewApp(ctx, nil, opt): %v\n", err)
//		return nil, err
//	}
//
//	return app.Firestore(ctx)
//}
//
////type PortfolioData struct {
//	Job         string
//	Description string
//	Timestamp   time.Time
//}
//
//func SaveItem(body string) error {
//	ctx := context.Background()
//
//	client, err := getFireBaseClient(ctx)
//	if err != nil {
//		fmt.Printf("Failed getFireBaseClient(ctx): %v\n", err)
//		return err
//	}
//
//	item := PortfolioData{
//		Job:         body,
//		Description: body,
//		Timestamp:   time.Now(),
//	}
//	_, err = client.Collection("portfolio-data-profile").Doc("1").Set(ctx, item)
//	if err != nil {
//		fmt.Printf("Failed client.Collection: %v\n", err)
//		return err
//	}
//	return nil
//}
