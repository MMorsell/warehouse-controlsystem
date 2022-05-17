import { Agent } from "./Agent.js"
import { moveAgent, clearAgentPath} from "./actions.js"
import { makeGrid } from "./grid.js";

const rows = 20;
const cols = 20;

// Used to track how many agents used this spot so we don't create holes in the path when clearing paths.
const refCount = new Array(cols * rows).fill(0);

let agents = new Map();

function advancePath(instruction) {
    const id = instruction.id;
    const x = instruction.x;
    const y = instruction.y;
    const index = XYtoLinear(x, y);

    if (!agents.has(id)) {
        agents.set(id, new Agent(id));
    }

    const agent = agents.get(id);

    moveAgent(agent, index);
    refCount[index]++;
}

function clearPath(id) {
    const agent = agents.get(id);
    clearAgentPath(agent, refCount);
}

function XYtoLinear(x, y) {
    return x + (y * cols);
}

makeGrid(rows, cols);

/*
advancePath({id: "1", x: 5, y: 2});
advancePath({id: "1", x: 5, y: 3});
advancePath({id: "1", x: 5, y: 4});
advancePath({id: "1", x: 5, y: 5});
advancePath({id: "1", x: 5, y: 5});
advancePath({id: "2", x: 2, y: 3});
advancePath({id: "2", x: 3, y: 3});
advancePath({id: "2", x: 4, y: 3});
advancePath({id: "2", x: 5, y: 3});
advancePath({id: "2", x: 6, y: 3});
advancePath({id: "2", x: 7, y: 3});
//clearPath("2")
*/

async function mock() {
    const paths = [
        {id: "1", x: 5, y: 7},
        {id: "2", x: 9, y: 5},
        {id: "1", x: 6, y: 7},
        {id: "2", x: 9, y: 6},
        {id: "1", x: 7, y: 7},
        {id: "2", x: 9, y: 7},
        {id: "1", x: 8, y: 7},
        {id: "2", x: 9, y: 8},
        {id: "1", x: 9, y: 7},
        {id: "2", x: 9, y: 9},
        {id: "1", x: 10, y: 7},
        {id: "2", x: 9, y: 10},
        {id: "1", x: 11, y: 7},
        {id: "2", x: 9, y: 11},
        {id: "1", x: 12, y: 7},
        {id: "2", x: 9, y: 12},
        {id: "1", x: 13, y: 7},
        {id: "2", x: 9, y: 13},
        {id: "1", x: 14, y: 7},
        {id: "2", x: 9, y: 14},
        {id: "1", x: 15, y: 7},
        {id: "2", x: 9, y: 15},
        {id: "1", x: 16, y: 7},
    ];

    for(let i = 0; i < paths.length - 1; i += 2) {
        advancePath(paths[i]);
        advancePath(paths[i+1]);
        await new Promise(r => setTimeout(r, 500));
    }
}

await mock()
clearPath("1")
await new Promise(r => setTimeout(r, 1000));
clearPath("2")



