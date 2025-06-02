<p align="center">
  <a href="#" target="blank"><img src="https://cdn-icons-png.flaticon.com/512/4712/4712027.png" alt="foodchat Logo" style="width: 250px; height: 155px;" width="250" height="155" /></a>
</p>

A modern, mvc-app using **Golang**, **OpenAI**, **MongoDB**, **Redis**, **Resend**, and **HTML**. This App helps users plan thier meals and gives them advice and nutritional tips.

---

## 🛠️ Getting Started

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd <project-directory>
```
### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Environment Setup

Copy the example environment file:

```bash
cp .env.example .env
```

Configure your `.env` file with appropriate values.

---

## 🚧 Development

Start the development server:

```bash
air
```

The server will restart automatically when files change.

---

## 🎯 Features

- **📤 Email Sending** - Send single emails at login for token verification.
- **🔑 Chat with OpenAI** - Create chat and ask OPENAI models questions about food
- **🌐 Chat Management** - Store chat in redis for 1 hour, allow user to backup for a lifetime
- **📎 File Attachments** - Support for binary data and remote URL attachments
- **👥 Profile Management** - Create, update profile and email

---

## People

- Author - [TylerJusfly](https://tylerjusfly.netlify.app/)
