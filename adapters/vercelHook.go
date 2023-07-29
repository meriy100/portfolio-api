package adapters

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func InitialVercelHookRequest(ctx context.Context) error {

	fmt.Println(os.Getenv("VERCEL_DEPLOY_HOOK_URL"))
	req, _ := http.NewRequestWithContext(ctx, "POST", os.Getenv("VERCEL_DEPLOY_HOOK_URL"), nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("vercel deploy hook request error. %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		return nil
	}
	return fmt.Errorf("vercel deploy hook request error")
}
