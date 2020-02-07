import matplotlib
matplotlib.use("Agg")
from matplotlib import pyplot as plt

import pandas as pd
import json

def label_point(x, y, val, ax):
    a = pd.concat({'x': x, 'y': y, 'val': val}, axis=1)
    for _, point in a.iterrows():
        ax.text(point['x'], point['y'], str(point['val']))

world = {}
parsed_citizens = {"Citizens":[]}

with open('world_dump.json', 'r') as f:
    world = json.load(f)

for citizen in world['BodyBags']:
    parsed_citizen = {}
    parsed_citizen['X'] = citizen['Coordinate']['X']
    parsed_citizen['Y'] = citizen['Coordinate']['Y']
    parsed_citizen['Z'] = 'grey'
    parsed_citizens['Citizens'].append(parsed_citizen)

for citizen in world['Citizens']:
    parsed_citizen = {}
    parsed_citizen['X'] = citizen['Coordinate']['X']
    parsed_citizen['Y'] = citizen['Coordinate']['Y']
    parsed_citizen['Z'] = 'red'
    parsed_citizens['Citizens'].append(parsed_citizen)

df = pd.DataFrame(parsed_citizens['Citizens'])
df.plot.scatter(x='X', y='Y', c=df.Z)

plt.savefig('world.png')

