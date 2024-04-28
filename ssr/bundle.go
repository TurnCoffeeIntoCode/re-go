package ssr

import (
	esbuild "github.com/evanw/esbuild/pkg/api"
)

func serverBundle() string {
	result := esbuild.Build(esbuild.BuildOptions{
		EntryPoints: []string{"./pages/serverEntry.tsx"},
		Bundle:      true,
		Write:       false,
		Outdir:      "./dist",
		Format:      esbuild.FormatESModule,
		Platform:    esbuild.PlatformBrowser,
		Target:      esbuild.ES2015,
		Banner: map[string]string{
			"js": textEncoderPolyfill + processPolyfill + consolePolyfill,
		},
		Loader: map[string]esbuild.Loader{
			".jsx": esbuild.LoaderJSX,
			".tsx": esbuild.LoaderTSX,
		},
	})
	script := string(result.OutputFiles[0].Contents)

	return script
}

func clientBundle() string {
	clientResult := esbuild.Build(esbuild.BuildOptions{
		EntryPoints: []string{"./pages/clientEntry.tsx"},
		Bundle:      true,
		Write:       false,
	})
	clientBundleString := string(clientResult.OutputFiles[0].Contents)
	return clientBundleString
}

// [Yaffle/TextEncoderTextDecoder.js](https://gist.github.com/Yaffle/5458286)
var textEncoderPolyfill = `function TextEncoder(){} TextEncoder.prototype.encode=function(string){var octets=[],length=string.length,i=0;while(i<length){var codePoint=string.codePointAt(i),c=0,bits=0;codePoint<=0x7F?(c=0,bits=0x00):codePoint<=0x7FF?(c=6,bits=0xC0):codePoint<=0xFFFF?(c=12,bits=0xE0):codePoint<=0x1FFFFF&&(c=18,bits=0xF0),octets.push(bits|(codePoint>>c)),c-=6;while(c>=0){octets.push(0x80|((codePoint>>c)&0x3F)),c-=6}i+=codePoint>=0x10000?2:1}return octets};function TextDecoder(){} TextDecoder.prototype.decode=function(octets){var string="",i=0;while(i<octets.length){var octet=octets[i],bytesNeeded=0,codePoint=0;octet<=0x7F?(bytesNeeded=0,codePoint=octet&0xFF):octet<=0xDF?(bytesNeeded=1,codePoint=octet&0x1F):octet<=0xEF?(bytesNeeded=2,codePoint=octet&0x0F):octet<=0xF4&&(bytesNeeded=3,codePoint=octet&0x07),octets.length-i-bytesNeeded>0?function(){for(var k=0;k<bytesNeeded;){octet=octets[i+k+1],codePoint=(codePoint<<6)|(octet&0x3F),k+=1}}():codePoint=0xFFFD,bytesNeeded=octets.length-i,string+=String.fromCodePoint(codePoint),i+=bytesNeeded+1}return string};`
var processPolyfill = `var process = {env: {NODE_ENV: "production"}};`
var consolePolyfill = `var console = {log: function(){}};`
