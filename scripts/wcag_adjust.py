import os
import json
import re

PAL_DIR = os.path.join('eui', 'themes', 'palettes')
STYLE_DIR = os.path.join('eui', 'themes', 'styles')

hsv_re = re.compile(r"^\s*([0-9.]+)\s*,\s*([0-9.]+)\s*,\s*([0-9.]+)\s*$")

# Adjust palette colors for better WCAG compliance
for fname in os.listdir(PAL_DIR):
    if not fname.endswith('.json'):
        continue
    path = os.path.join(PAL_DIR, fname)
    with open(path, 'r') as f:
        data = json.load(f)
    colors = data.get('Colors', {})
    bg = colors.get('background')
    if bg:
        m = hsv_re.match(bg)
        if m:
            h, s, v = map(float, m.groups())
            # Choose text color based on background brightness
            colors['text'] = '0,0,1' if v < 0.5 else '0,0,0'
    acc = colors.get('accent')
    if acc:
        m = hsv_re.match(acc)
        if m:
            h, s, v = map(float, m.groups())
            v = max(0.6, v)
            s = max(0.5, s)
            colors['accent'] = f"{h},{s},{v}"
    data['Colors'] = colors
    with open(path, 'w') as f:
        json.dump(data, f, indent=2)

# Adjust styles - ensure visible borders
for fname in os.listdir(STYLE_DIR):
    if not fname.endswith('.json'):
        continue
    path = os.path.join(STYLE_DIR, fname)
    with open(path, 'r') as f:
        data = json.load(f)
    border = data.get('Border', {})
    changed = False
    for k, v in border.items():
        if v < 1:
            border[k] = 1
            changed = True
    if changed:
        data['Border'] = border
        with open(path, 'w') as f:
            json.dump(data, f, indent=2)
