f = open("quran.csv", "r")
f.readline()
for line in f:
    line = line.strip()
    if (line == ""):
        continue
    items = line.split(",")
    if (len(items) < 4):
        continue
    sura_id = int(items[1])
    verse_id = int(items[2])
    ayah = str(items[3]).strip('"')
    print(sura_id, verse_id, ayah)
