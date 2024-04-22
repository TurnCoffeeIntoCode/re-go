package ssr

import (
	"errors"
	"fmt"
	"log"

	v8 "rogchap.com/v8go"
)

var (
	iso                = v8.NewIsolate()
	script             *v8.UnboundScript
	serverBundleScript string
	clientBundleScript string
)

// Bundle the scripts once at startup
func init() {
	fmt.Println("Bundling scripts")
	serverBundleScript = serverBundle()
	clientBundleScript = clientBundle()
	fmt.Println("Initializing V8")
	var err error
	script, err = iso.CompileUnboundScript(serverBundleScript, "bundle.js", v8.CompileOptions{})
	if err != nil {
		log.Fatalf("Failed to compile bundled script: %v", err)
	}
}

// Render the HTML using V8
func renderHTML(props string) (string, error) {
	// Create a new context
	ctx := v8.NewContext(iso)
	defer ctx.Close()

	// Run the pre bundled script
	_, err := script.Run(ctx)
	if err != nil {
		log.Fatalf("Failed to run bundled script: %v", err)
		return "", errors.New("failed to run bundled script")
	}

	// Pass props to the renderApp function
	val, err := ctx.RunScript(fmt.Sprintf("renderApp(%s)", props), "render.js")
	if err != nil {
		log.Fatalf("Failed to render React component: %v", err)
		return "", errors.New("failed to render React component")
	}

	return val.String(), nil
}
