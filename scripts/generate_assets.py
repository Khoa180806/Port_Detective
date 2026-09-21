import os
from PIL import Image, ImageDraw, ImageFont

BG_COLOR = (18, 22, 32)
HEADER_COLOR = (26, 32, 46)
BORDER_COLOR = (42, 50, 68)
TEXT_WHITE = (248, 250, 252)
TEXT_GRAY = (148, 163, 184)
TEXT_DARK_GRAY = (90, 105, 125)
TEXT_CYAN = (56, 189, 248)
TEXT_GREEN = (74, 222, 128)
TEXT_YELLOW = (250, 204, 21)
TEXT_RED = (248, 113, 113)

FONT_PATH = "C:/Windows/Fonts/consola.ttf"
FONT_SIZE = 16
font = ImageFont.truetype(FONT_PATH, FONT_SIZE)
font_bold = ImageFont.truetype("C:/Windows/Fonts/consolab.ttf", FONT_SIZE) if os.path.exists("C:/Windows/Fonts/consolab.ttf") else font
font_header = ImageFont.truetype("C:/Windows/Fonts/arial.ttf", 13)

def draw_window(width, height, title):
    img = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    draw = ImageDraw.Draw(img)

    # Soft drop shadow
    for i in range(12, 0, -1):
        alpha = int(32 * (1 - i / 12))
        draw.rounded_rectangle([i, i, width - i, height - i], radius=14, fill=(0, 0, 0, alpha))

    # Window frame
    rect = [10, 10, width - 10, height - 10]
    draw.rounded_rectangle(rect, radius=10, fill=BG_COLOR, outline=BORDER_COLOR, width=1)

    # Window header
    draw.rounded_rectangle([10, 10, width - 10, 44], radius=10, fill=HEADER_COLOR)
    draw.rectangle([10, 36, width - 10, 44], fill=HEADER_COLOR)
    draw.line([10, 44, width - 10, 44], fill=BORDER_COLOR, width=1)

    # Mac-style buttons
    draw.ellipse([24, 22, 36, 34], fill=(255, 95, 86))
    draw.ellipse([44, 22, 56, 34], fill=(255, 189, 46))
    draw.ellipse([64, 22, 76, 34], fill=(255, 189, 46))
    draw.ellipse([64, 22, 76, 34], fill=(39, 201, 63))

    # Centered Header Title
    title_w = draw.textlength(title, font=font_header)
    draw.text(((width - title_w) // 2, 19), title, font=font_header, fill=(148, 163, 184))

    return img, draw

# ================= IMAGE 1: CHECK & KILL DEMO =================
def render_check():
    width, height = 760, 380
    img, draw = draw_window(width, height, "port-detective check & kill")

    y = 64
    x = 30
    lh = 26

    # Command 1: pd check 8080
    draw.text((x, y), "$ ", font=font_bold, fill=TEXT_GREEN)
    draw.text((x + 18, y), "pd check 8080", font=font_bold, fill=TEXT_WHITE)
    y += lh + 2

    # Output: Port 8080 is occupied by:
    draw.text((x, y), "Port ", font=font, fill=TEXT_WHITE)
    draw.text((x + 44, y), "8080", font=font_bold, fill=TEXT_CYAN)
    draw.text((x + 82, y), " is occupied by:", font=font, fill=TEXT_WHITE)
    y += lh

    # Details
    draw.text((x + 18, y), "PID:      ", font=font, fill=TEXT_GRAY)
    draw.text((x + 104, y), "14280", font=font_bold, fill=TEXT_RED)
    y += lh

    draw.text((x + 18, y), "Process:  ", font=font, fill=TEXT_GRAY)
    draw.text((x + 104, y), "node.exe", font=font, fill=TEXT_YELLOW)
    y += lh

    draw.text((x + 18, y), "Command:  ", font=font, fill=TEXT_GRAY)
    draw.text((x + 104, y), "node server.js", font=font, fill=TEXT_DARK_GRAY)
    y += lh

    draw.text((x + 18, y), "Protocol: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 104, y), "tcp", font=font_bold, fill=TEXT_GREEN)
    y += lh + 14

    # Command 2: pd kill 8080 --force
    draw.text((x, y), "$ ", font=font_bold, fill=TEXT_GREEN)
    draw.text((x + 18, y), "pd kill 8080 --force", font=font_bold, fill=TEXT_WHITE)
    y += lh + 2

    draw.text((x, y), "Terminating PID 14280 (node.exe)... ", font=font, fill=TEXT_GRAY)
    draw.text((x + 315, y), "SUCCESS", font=font_bold, fill=TEXT_GREEN)

    os.makedirs("docs/images", exist_ok=True)
    img.save("docs/images/check_demo.png")
    print("Regenerated check_demo.png")

# ================= IMAGE 2: SCAN DEMO =================
def render_scan():
    width, height = 760, 430
    img, draw = draw_window(width, height, "port-detective scan")

    y = 64
    x = 30
    lh = 24

    draw.text((x, y), "$ ", font=font_bold, fill=TEXT_GREEN)
    draw.text((x + 18, y), "pd scan 3000-3005", font=font_bold, fill=TEXT_WHITE)
    y += lh + 2

    draw.text((x, y), "Scanning ports from 3000 to 3005...", font=font, fill=TEXT_GRAY)
    y += lh

    draw.text((x, y), "Found 3 processes:", font=font_bold, fill=TEXT_WHITE)
    y += lh + 4

    # Process 1
    draw.text((x + 16, y), "Port: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 64, y), "3000", font=font_bold, fill=TEXT_CYAN)
    draw.text((x + 115, y), "|  PID: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 180, y), "18240", font=font_bold, fill=TEXT_RED)
    draw.text((x + 240, y), "|  Process: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 340, y), "vite.exe", font=font, fill=TEXT_YELLOW)
    draw.text((x + 470, y), "|  Protocol: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 580, y), "tcp", font=font_bold, fill=TEXT_GREEN)
    y += lh

    draw.text((x + 16, y), "----------------------------------------------------------------------", font=font, fill=TEXT_DARK_GRAY)
    y += lh

    # Process 2
    draw.text((x + 16, y), "Port: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 64, y), "3001", font=font_bold, fill=TEXT_CYAN)
    draw.text((x + 115, y), "|  PID: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 180, y), "9412", font=font_bold, fill=TEXT_RED)
    draw.text((x + 240, y), "|  Process: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 340, y), "react-dev.exe", font=font, fill=TEXT_YELLOW)
    draw.text((x + 470, y), "|  Protocol: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 580, y), "tcp", font=font_bold, fill=TEXT_GREEN)
    y += lh

    draw.text((x + 16, y), "----------------------------------------------------------------------", font=font, fill=TEXT_DARK_GRAY)
    y += lh

    # Process 3
    draw.text((x + 16, y), "Port: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 64, y), "3004", font=font_bold, fill=TEXT_CYAN)
    draw.text((x + 115, y), "|  PID: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 180, y), "23150", font=font_bold, fill=TEXT_RED)
    draw.text((x + 240, y), "|  Process: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 340, y), "docker-proxy", font=font, fill=TEXT_YELLOW)
    draw.text((x + 470, y), "|  Protocol: ", font=font, fill=TEXT_GRAY)
    draw.text((x + 580, y), "tcp", font=font_bold, fill=TEXT_GREEN)
    y += lh + 14

    draw.text((x, y), "$ ", font=font_bold, fill=TEXT_GREEN)
    draw.rectangle([x + 18, y + 2, x + 27, y + 18], fill=TEXT_CYAN)

    img.save("docs/images/scan_demo.png")
    print("Regenerated scan_demo.png")

# ================= GIF: INTERACTIVE TERMINAL ANIMATION =================
def render_gif():
    width, height = 760, 360
    frames = []

    # Sequence of typing
    text_to_type = "pd check 8080"
    for i in range(1, len(text_to_type) + 1):
        img, draw = draw_window(width, height, "port-detective (interactive)")
        draw.text((30, 64), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, 64), text_to_type[:i], font=font_bold, fill=TEXT_WHITE)
        cx = 48 + int(draw.textlength(text_to_type[:i], font=font_bold))
        draw.rectangle([cx + 2, 66, cx + 11, 82], fill=TEXT_CYAN)
        frames.append((img.convert("RGB"), 110))

    # Pause after typing
    for _ in range(3):
        img, draw = draw_window(width, height, "port-detective (interactive)")
        draw.text((30, 64), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, 64), text_to_type, font=font_bold, fill=TEXT_WHITE)
        frames.append((img.convert("RGB"), 130))

    # Step-by-step output lines
    def draw_first_output(draw):
        y = 90
        lh = 25
        draw.text((30, y), "Port ", font=font, fill=TEXT_WHITE)
        draw.text((30 + 44, y), "8080", font=font_bold, fill=TEXT_CYAN)
        draw.text((30 + 82, y), " is occupied by:", font=font, fill=TEXT_WHITE)
        y += lh
        draw.text((48, y), "PID:      ", font=font, fill=TEXT_GRAY)
        draw.text((134, y), "14280", font=font_bold, fill=TEXT_RED)
        y += lh
        draw.text((48, y), "Process:  ", font=font, fill=TEXT_GRAY)
        draw.text((134, y), "node.exe", font=font, fill=TEXT_YELLOW)
        y += lh
        draw.text((48, y), "Command:  ", font=font, fill=TEXT_GRAY)
        draw.text((134, y), "node server.js", font=font, fill=TEXT_DARK_GRAY)
        y += lh
        draw.text((48, y), "Protocol: ", font=font, fill=TEXT_GRAY)
        draw.text((134, y), "tcp", font=font_bold, fill=TEXT_GREEN)

    # Frame with check output
    for _ in range(6):
        img, draw = draw_window(width, height, "port-detective (interactive)")
        draw.text((30, 64), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, 64), text_to_type, font=font_bold, fill=TEXT_WHITE)
        draw_first_output(draw)
        frames.append((img.convert("RGB"), 150))

    # Typing kill command
    kill_text = "pd kill 8080 --force"
    for i in range(1, len(kill_text) + 1, 2):
        img, draw = draw_window(width, height, "port-detective (interactive)")
        draw.text((30, 64), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, 64), text_to_type, font=font_bold, fill=TEXT_WHITE)
        draw_first_output(draw)

        y = 226
        draw.text((30, y), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, y), kill_text[:i], font=font_bold, fill=TEXT_WHITE)
        cx = 48 + int(draw.textlength(kill_text[:i], font=font_bold))
        draw.rectangle([cx + 2, y + 2, cx + 11, y + 18], fill=TEXT_CYAN)
        frames.append((img.convert("RGB"), 80))

    # Success frame
    for _ in range(16):
        img, draw = draw_window(width, height, "port-detective (interactive)")
        draw.text((30, 64), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, 64), text_to_type, font=font_bold, fill=TEXT_WHITE)
        draw_first_output(draw)

        y = 226
        draw.text((30, y), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.text((48, y), kill_text, font=font_bold, fill=TEXT_WHITE)
        y += 28

        draw.text((30, y), "Terminating PID 14280 (node.exe)... ", font=font, fill=TEXT_GRAY)
        draw.text((345, y), "SUCCESS", font=font_bold, fill=TEXT_GREEN)
        y += 30

        draw.text((30, y), "$ ", font=font_bold, fill=TEXT_GREEN)
        draw.rectangle([48, y + 2, 57, y + 18], fill=TEXT_CYAN)
        frames.append((img.convert("RGB"), 180))

    # Save animated GIF
    imgs = [f[0] for f in frames]
    durations = [f[1] for f in frames]
    imgs[0].save(
        "docs/images/demo.gif",
        save_all=True,
        append_images=imgs[1:],
        duration=durations,
        loop=0
    )
    print("Regenerated demo.gif")

if __name__ == "__main__":
    render_check()
    render_scan()
    render_gif()
