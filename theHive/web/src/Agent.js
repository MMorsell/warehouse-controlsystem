export class Agent {
    constructor(id) {
        this.id = id;
        this.path = [];
    }

    updatePath(index) {
        this.path.push(index)
    }

    clearPath() {
        this.path = []
    }
}




