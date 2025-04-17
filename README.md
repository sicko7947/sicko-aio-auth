# 🔐 sicko-aio-auth

`sicko-aio-auth` is the authentication service for the **SickoAIO bot framework**.  
It handles secure user validation, machine verification, and webhook communication.

---

## 🧠 What It Does

- ✅ **Global Success Webhook Handling**  
  Sends webhook events (e.g., successful checkouts) to configured endpoints globally.

- 🔁 **gRPC Auth (Bidirectional Stream)**  
  Constantly verifies:
  - Machine ID  
  - IP address  
  Helps prevent cracks, spoofing, and unauthorized bot usage.

- 🗂️ **User Database Handler**  
  Handles user license data and authentication states (simple implementation, can be swapped for real DB).

---

## 🚀 How to Run

```bash
go run main.go <port>
