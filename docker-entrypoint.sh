#!/bin/sh
set -e

# Проверяем токен окружения
if [ -z "$TELEGRAM_BOT_TOKEN" ]; then
  echo "Error: TELEGRAM_BOT_TOKEN environment variable is not set." >&2
  exit 1
fi

# Запускаем бота из скомпилированного образа
exec ./main