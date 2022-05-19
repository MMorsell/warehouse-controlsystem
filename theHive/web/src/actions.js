const cells = document.getElementById("grid").children

export function moveAgent(agent, index) {
    agent.updatePath(index);
    draw(agent, index);
}

function draw(agent, index) {
    // Mark previous position as used.
    if (agent.path === undefined || agent.path.length > 1) {
        const prevIndex = agent.path.at(-2);
        cells.item(prevIndex).className = "cellUsed";
    }

    // New current position.
    cells.item(index).className = "cellCurrent";
}

export function clearAgentPath(agent, refCount) {
    for(let i = 0; i < agent.path.length; i++) {
        var index = agent.path[i]; // function scope so we can use after to fix head
        if (refCount[index] == 1) {
            cells.item(index).className = "cellFree";
            
        }
        refCount[index]--;
    }

    // Correction if last agent's final position was on a used one.
    if (refCount[index] !== 0) {
        cells.item(index).className = "cellUsed";
    }
    agent.clearPath();
}