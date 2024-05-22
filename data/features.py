import pandas as pd

# Need to run this script from the `data` directory

df = pd.read_csv("./input/megaGymDataset.csv", header=0, index_col=0)

print("Type:", df['Type'].unique())
print("BodyPart:", df['BodyPart'].unique())
print("Equipment:", df['Equipment'].unique())
print("Level:", df['Level'].unique())

# Equipment has a "None" column that Python reads as the None type
df['Equipment'].fillna("None", inplace=True) 
vectors = {}

def setupVectors():
  types = df["Type"].unique()
  for type in types:
    vectors[type] = []

  bps = df["BodyPart"].unique()
  for bp in bps:
    vectors[bp] = []
  
  eqs = df['Equipment'].unique()
  for eq in eqs:
    vectors[eq] = []

  vectors["Level"] = 0

def addType(type):
  types = df['Type'].unique()
  for t in types:
    if t == type:
      vectors[t].append(1)
    else:
      vectors[t].append(0)

def addBodypart(part):
  bps = df['BodyPart'].unique()
  for bp in bps:
    if bp == part:
      vectors[bp].append(1)
    else:
      vectors[bp].append(0)

def addEquipment(equipment):
  eqs = df['Equipment'].unique()
  for e in eqs:
    if e == equipment:
      vectors[e].append(1)
    else:
      vectors[e].append(0)

def addLevel(level):
  levels = {
    "Expert": 2,
    "Intermediate": 1,
    "Beginner": 0,
  }

  vectors["Level"] = levels[level]

setupVectors()
for index, row in df.iterrows():
  addType(row["Type"])
  addBodypart(row["BodyPart"])
  addEquipment(row["Equipment"])
  addLevel(row["Level"])

vector_df = pd.DataFrame(vectors)

print(vector_df.head())
print(vector_df.tail())
print(vector_df.loc[932])

vector_df.to_csv("./output/features.csv")