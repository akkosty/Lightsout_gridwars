# bot/main.py
import logging
from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup
from telegram.ext import (
    ApplicationBuilder,
    CommandHandler,
    ContextTypes,
    CallbackQueryHandler,
)

# -------------------------------------------------
# Конфигурация логгера
logging.basicConfig(
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
    level=logging.INFO,
)
logger = logging.getLogger(__name__)

# -------------------------------------------------
# Токен берём из переменной окружения (без hard‑code!)
import os

TOKEN = os.getenv("TELEGRAM_BOT_TOKEN")
if not TOKEN:
    raise RuntimeError("Установите переменную окружения TELEGRAM_BOT_TOKEN")

# -------------------------------------------------
async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    """Ответ на /start"""
    keyboard = [
        [InlineKeyboardButton("👋 Привет, Танечка! Ты лучшая ❤️", callback_data="say_hello")],
        [InlineKeyboardButton("ℹ️ Инфо о боте", callback_data="info")],
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await update.message.reply_text(
        "Привет! Я базовый скелет Telegram‑бота.\nВыбери действие:",
        reply_markup=reply_markup,
    )

# -------------------------------------------------
async def button_handler(update: Update, context: ContextTypes.DEFAULT_TYPE):
    """Обработчик нажатий inline‑кнопок"""
    query = update.callback_query
    await query.answer()  # обязательный ACK

    if query.data == "say_hello":
        await query.edit_message_text("👋 Привет от бота!")
    elif query.data == "info":
        await query.edit_message_text(
            "🤖 Я простой бот‑скелет. Добавляйте свои команды и логику!"
        )
    else:
        await query.edit_message_text(f"❓ Неизвестный callback: {query.data}")

# -------------------------------------------------
async def echo(update: Update, context: ContextTypes.DEFAULT_TYPE):
    """Эхо‑команда для отладки"""
    if update.message:
        await update.message.reply_text(f"🗣 Вы сказали: {update.message.text}")

# -------------------------------------------------
def main() -> None:
    app = ApplicationBuilder().token(TOKEN).build()

    # Регистрация хэндлеров
    app.add_handler(CommandHandler("start", start))
    app.add_handler(CallbackQueryHandler(button_handler))
    app.add_handler(CommandHandler("echo", echo))

    # Запуск в режиме polling (для локального теста)
    logger.info("Бот запущен, начинаем polling…")
    app.run_polling()

if __name__ == "__main__":
    main()
