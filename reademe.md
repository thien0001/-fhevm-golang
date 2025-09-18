# FHEVM Go Demo 🚀

Demo application **Zama FHEVM** written in **Golang + Fiber**.
This project demonstrates the use of **Fully Homomorphic Encryption (FHE)** on an EVM-compatible blockchain (Sepolia testnet).

---

## ✨ Features
- Input data encryption API (`/api/encrypt`)
- Send FHE transactions on-chain (`/api/send`)
- Get results from blockchain (`/api/result/:txHash`)
- Written in **Go Fiber** – fast, lightweight
- Prepared for UI integration (React / Next.js)

---

## ⚡ Installation

### 1. Clone repo
```bash
git clone https://github.com/thien0001/-fhevm-golang
cd fhevm-go-demo


Create a .env file in the root directory:

RPC_URL=https://sepolia.infura.io/v3/<YOUR_INFURA_PROJECT_ID>
PRIVATE_KEY=<YOUR_WALLET_PRIVATE_KEY>

🔑 Note:

PRIVATE_KEY must be 64-character hex, without prefix 0x.

RPC_URL can use Infura, Alchemy or other provider.

3. Install dependencies
go mod tidy

4. Run the server
go run ./cmd/api

The server will run at:
👉 http://localhost:3000

🔌 API Endpoints
POST /api/encrypt

Encrypt the input data.

Request:

{
"data": "42"
}

Response:

{
"encrypted": "0xabcd1234..."
}

POST /api/send

Send encrypted data to the blockchain.

Request:

{
"encrypted": "0xabcd1234..."
}

Response:

{
"txHash": "0x1234abcd..."
}

GET /api/result/:txHash

Get the transaction result from the blockchain.

Response:

{
"result": "42"
}

🛠️ Roadmap

Frontend integration (test html)

Private DeFi demo (add/subtract on encrypted data)

Private Identity demo (verify without exposing data)

🌐 Contact

Demo built for Zama community
.
All contributions and pull requests are welcome!
