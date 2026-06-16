# Базовый образ
FROM python:3.11-slim

# Устанавливаем зависимости для работы Telegram Bot API (необязательно, но полезно)
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc libpq-dev && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .

# Установка зависимостей
RUN pip install --upgrade pip && \
    pip install -r requirements.txt

# Переменная окружения, которую будет подставлять CI/CD (не храните токен в Dockerfile!)
ENV TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}

CMD ["python", "bot/main.py"]
