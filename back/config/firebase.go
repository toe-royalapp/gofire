package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {
	ctx := context.Background()
	// serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	// if err != nil {
	// 	panic("Unable to load serviceAccountKeys.json file")
	// }

	// opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	opt := option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "gofire-9ce2c",
		"private_key_id": "1c84d7d26fb54c8827a47dfec175b0dcb6821a3d",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC05eeIzpMgz9Ke\nuMTylXtxK8xO4n5oSXbFUhcZjiQA+ioTvHS6M7mSMnZJEXkCHkSVEvV1OCrxQ43w\nJYOsGO4Be0mBCSyNQa8h7pez1LPIJvQuRLIn1KWhxFFzBZHVgJuAFDSXdiRrON48\nOQQyNrvGWtk79T4/0C77U5RmxZEIHOBZ8upl8yqJEMP4BJ1FQ56PZuueX1iK/wNp\ntVcUx/PP30c1THybbpE/0dwrD1YdGdCh9cpPbNq3utEO1+CWJD3VwVxo/O4hcHUJ\navoM+bm0X7EGPQjorHqBnJ0NDDYf/VnN9DNP0UHySgQurEWoNXBtKwk05TjYgMWI\nj3NmnEhDAgMBAAECggEAJjsSaEDUKlKRnjbJdMUvvfa2KpPGiE7CWahkNmPqawRz\nwREKkFGfLe0ZP48ARnjUILdNhdT1imCvDio3fjcUkd6W5bzlHANeOmscx4Yz6qCE\n6YReLHnN+XauigMK05bnBjX++WWgA8MUgKDCKMSVgbGwiHHnkn/yde4vhlrJsG9i\nytBQ22okHbL9zCHmKsH5n18RZaNX7lSckDNEcLvRrxTCXQbuVDwhLIkYs5cJH1d3\ntV305qIGwD1Xywl0KBKULv1aPSQSmAmm+d5xyTMAolI0BQplEa/kDmpXucQAqnFx\nCq/dtBp75caPEwZKkvDdXFVWRt5nJyAQ1S6W5IUAZQKBgQDizZAlsJzH2WTnF4JT\nO5a5qdoJt6W5J9hKUlJSIQEKBWo9EUesFmEzdqZFItOpotnMRIzcnnzgCyCVA78N\nt4AfObaP+RAsRDxj4knvm2gu1Em/0mpQgEsguWXpy1MUjOTDCijMQRJLryL9zXGE\nnSIbetdVVoHEMq0iiGMLiEfaDQKBgQDML4P6tjANDpk2o0ezl4CxrHPZ8yHNHtii\nSWVa19ghu1RvvKfr0090Te3l9fLt+97LiqwCBVncmOqPuk2GT/AlH6HBRgbz5VAb\nCBtO5KExBee1rrDnTD2Se4d6WsN25xvjcIVMWaBv52mV1bGlC9y3zLqMuhpWDsdk\nd8KuXGenjwKBgFngolX9RjQAV3contHDFHjg8XHWYAse2hyhwNOhFptVCAPJPEDa\n4YwWYc+V/JEF4w+KvtOSzuOuJSxIPsb9x/0XztwBFEKmi9P9UdVtHX0pTUyB4vWh\n0aPXNKbQl0zWhLUx6nb+9nQdpF01s92cs252YK0FygjpGClOKQnh6K9lAoGAKyQy\n94r89hVKi//Ny4VMPL4aMEetsaA913Q4hQwr71ycR4uN00bvd7xrcnYX4O6hsVHY\nbAjjZresf7e2X5WOQNnBHIwGNgwSc9OvqWinIiuEPJ/vYr96FvQguNpqiDnma5S/\npVUej+ZOKclf4mZyOSd5lvhPTjWLLZAcQyjlAcMCgYEAq26VPiSgQ2d9hJwom4WS\ncL7veXQONxAvQoJKX0T3P3gKYb/818GRYVEnW45lktS/9KkoO/Oa3wevD/Qoje9p\nN6jFP+NOrVcLI9RztRQa/Xby1LGNW3OG1ak4vC4U6F07DIRaSrdacoE+0Tgl/V3N\nYhMKFDRGHRWd34cYuzW28P8=\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-nhbpg@gofire-9ce2c.iam.gserviceaccount.com",
		"client_id": "113864022722627822667",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-nhbpg%40gofire-9ce2c.iam.gserviceaccount.com"
	}`))

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	client, _ := app.Messaging(ctx)

	return app, ctx, client
}
