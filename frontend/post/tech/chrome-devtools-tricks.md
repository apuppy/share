# Chrome DevTools Tricks

## base64 decode
```javascript
// atob(): decodes a Base64-encoded string ("atob" should be read as "ASCII to binary").

atob('SGVsbG8gd29ybGQh')
>'Hello world!'
```

## base64 encode
```javascript
// btoa(): creates a Base64-encoded ASCII string from a "string" of binary data ("btoa" should be read as "binary to ASCII").

btoa("Hello world!")
>'SGVsbG8gd29ybGQh'
```

## example
```javascript
const encodedData = btoa("Hello, world"); // encode a string
const decodedData = atob(encodedData); // decode the string
```

## work with URL encode/decode
```javascript
function utf8_to_b64(str) {
  return window.btoa(unescape(encodeURIComponent(str)));
}

function b64_to_utf8(str) {
  return decodeURIComponent(escape(window.atob(str)));
}

// Usage:
utf8_to_b64("Hello world !"); // "SGVsbG8gd29ybGQgIQ=="
b64_to_utf8("SGVsbG8gd29ybGQgIQ=="); // "Hello world !"
```

## reference
[Base64 - MDN Web Docs Glossary: Definitions of Web-related terms | MDN](https://developer.mozilla.org/en-US/docs/Glossary/Base64)