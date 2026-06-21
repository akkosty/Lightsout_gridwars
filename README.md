# 🦄 Lightsout GridWars Telegram Bot (Go Edition)!!!

Это минимальный шаблон Go бота для Telegram с использованием библиотеки `go-telegram/bot`. 
Он готов к запуску в Docker и CI/CD через GitHub Actions. !!!

## 🚀 Как запустить локально!!!

1. **Создайте бота** у @BotFather → получите токен!!!
2. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/your‑login/your‑repo.git
   cd your-repo
   ```

3. Установите зависимости и соберите проект:

   ```bash
   go mod download
   go build -o main ./bot
   ```

4. Создайте `.env` файл с токеном бота и запустите:

   ```bash
   TELEGRAM_BOT_TOKEN=ваш_токен DEBUG=True ./main
   ```

## 🐳 Как запустить в Docker!!!

```bash
docker-compose up -d
```

## ⚙️ Настройки через переменные окружения!!!

Создайте файл `.env`:

```
TELEGRAM_BOT_TOKEN=ваш_токен_от_BotFather
DEBUG=True
```

---

### 🛠 GitHub Actions!!!

Этот репозиторий поддерживает автоматический CI/CD через GitHub Actions. 

### ⚠️ Важные примечания!!!

- **Безопасность:** Никогда не храните токены в коде!!!  
- **Тестирование:** Тестируйте бота локально перед деплоем!!!  
- **Деплой:** Используйте `render.yaml` для автоматического развёртывания на Render!!!