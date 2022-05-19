import { Agent } from "./Agent.js"
import { moveAgent, clearAgentPath} from "./actions.js"
import { makeGrid } from "./grid.js";

const rows = 20;
const cols = 20;

// Used to track how many agents used this spot so we don't create holes in the path when clearing paths.
const refCount = new Array(cols * rows).fill(0);

let agents = new Map();

export function advancePath(instruction) {
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

export function clearPath(id) {
    const agent = agents.get(id);
    clearAgentPath(agent, refCount);
}

function XYtoLinear(x, y) {
    const linearVersion = x + (y * cols);
    console.debug(`converting x '${x}' and '${y}' to linear value '${linearVersion}'`)
    return linearVersion
}

makeGrid(rows, cols);