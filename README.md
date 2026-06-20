# 📦 Базовый скелет Telegram‑бота !!!

Это минимальный шаблон бота на Python + `python-telegram-bot`. 
Он уже готов к запуску в Docker и CI/CD через GitHub Actions. !!!

## 🚀 Как запустить локально (Polling) !!!

1. **Создайте бота** у @BotFather → получите токен !!!
2. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/your‑login/your‑repo.git
   cd your-repo
   ```

3. Установите зависимости:

   ```bash
   pip install -r requirements.txt
   ```

4. Создайте `.env` файл с токеном бота и запустите:

   ```bash
   python bot.py
   ```

## 🐳 Как запустить в Docker !!!

```bash
docker-compose up -d
```

## ⚙️ Настройки через переменные окружения !!!

Создайте файл `.env`:

```
BOT_TOKEN=ваш_токен_от_BotFather
DEBUG=True
```

---

### 🛠 GitHub Actions !!!

Этот репозиторий поддерживает автоматический CI/CD через GitHub Actions. 

### ⚠️ Важные примечания !!!

- **Безопасность:** Никогда не храните токены в коде!!!  
- **Тестирование:** Тестируйте бота локально перед деплоем!!!  
- **Деплой:** Используйте `render.yaml` для автоматического развёртывания на Render!!!