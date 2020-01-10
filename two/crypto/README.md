# Practice for golang rsa private key encrypt and public key decrypt

1. use encrypt.go to encrypt `text` with key and write into file
2. read file and use decrypt.go to decrypt `file content`

# 為什麼要做crytpo
一些具有交易隱密性的文件傳輸，為了確保文件能保密，在傳輸前預先做加密。e.g. :貨幣交易，個人隱密資料


# Hashing Passwords to Compatiable Cipher Keys
When encrypting and decrypting data, it is important that you are using a 32 character, or 32 byte key.

Being realistic, you’re probably going to want to use a passphrase and that passphrase will never be 32 characters in length.

To get around this, you can actually hash your passphrase using a hashing algorithm that produces 32 byte hashes.

I found a list of hashing algorithms on Wikipedia that provide output lengths.

We’re going to be using a simple MD5 hash. It is insecure, but it doesn’t really matter since we won’t be storing the output.

Within a Go project, we can add the following function:

```go
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
```
The above function will take a passphrase or any string, hash it, then return the hash as a hexadecimal value.

Remember, we just need keys that meet the length criteria that AES demands.

- 使用MD5驗證一些字串的吻合，而非使用字串長度或者亂數來做驗證。此驗證模式的key為一個32位元長度的加密key

# Encrypting Data with an AES Cipher

Now that we have a key of an appropriate size, we can start the encryption process.

We can be encrypting text, or any binary data, it doesn’t really matter.

Within a Go project, include the following function:

```go
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
```
Remember, the code I’m using is a Frankenstein from several different sources. So what exactly is happening in the above function?

```go
block, _ := aes.NewCipher([]byte(createHash(passphrase)))
```
First we create a new block cipher based on the hashed passphrase.

Once we have our block cipher, we want to wrap it in Galois Counter Mode (GCM) with a standard nonce length.

Before we can create the ciphertext, we need to create a nonce.

```go
nonce := make([]byte, gcm.NonceSize())
if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
    panic(err.Error())
}
```

The nonce that we create needs to be the length specified by GCM.

It is important to note that the nonce used for decryption must be the same nonce used for encryption.

- There are a few strategies that can be used to make sure our decryption nonce matches the encryption nonce.

1. One strategy would be to store the nonce alongside the encrypted data if it is going into a database. 

2. Another option is to prepend or append the nonce to the encrypted data. We’ll prepending the nonce.

```go
ciphertext := gcm.Seal(nonce, nonce, data, nil)
```

The first parameter in the `Seal` command is our prefix value.

The encrypted data will be appended to it. With the ciphertext, we can return it back to a calling function.


# Decrypting Data that uses an AES Cipher

Now that we have potentially encrypted some data, we probably want to be sure that we can decrypt that same data.

The process for decryption is nearly the same as the encryption process.

```go
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
```

In the above code we create a new block cipher using a hashed passphrase.

We wrap the block cipher in Galois Counter Mode and get the nonce size.

Remember, we prefixed our encrypted data with the nonce.

This means that we need to separate the nonce and the encrypted data.

```go
nonceSize := gcm.NonceSize()
nonce, ciphertext := data[:nonceSize], data[nonceSize:]
```

When we have our nonce and ciphertext separated, we can decrypt the data and return it as plaintext.

# Encrypting and Decrypting Files

Encrypting and decrypting data as we demand is cool, but what if we wanted to encrypt and decrypt files.

One way to handle file encryption is to take what we’ve already done and just use the file commands.

```go
func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}
```

The above function wile create and open a file based on the filename passed.

With the file open, we can encrypt some data and write it to the file.

The file will close when we’re done.

To decrypt the file, we could use the following function:

```go
func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}
```

Now I’m not saying that this is the best way to handle file encryption and decryption, 

but I’m saying that this is just one of many ways to accomplish the task.


# 小結:

用hash做資料校驗比對，通常用MD5

用 encrypt/decrypt 做 加/解密 來解開封包內容


# refer:
- https://gist.github.com/kkHAIKE/be3b8d7ff8886457c6fdac2714d56fe1
- https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/