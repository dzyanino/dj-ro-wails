export namespace types {
	
	export class FixablePosition {
	    x: number;
	    y: number;
	    fixed?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FixablePosition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.fixed = source["fixed"];
	    }
	}
	export class Layouts {
	    nodes: Record<string, FixablePosition>;
	
	    static createFrom(source: any = {}) {
	        return new Layouts(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nodes = this.convertValues(source["nodes"], FixablePosition, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NodeProps {
	    weightTo: number;
	    previousNode: string;
	    marked: boolean;
	    valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NodeProps(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.weightTo = source["weightTo"];
	        this.previousNode = source["previousNode"];
	        this.marked = source["marked"];
	        this.valid = source["valid"];
	    }
	}
	export class ResolutionNode {
	    name?: string;
	    properties?: Record<string, any>;
	    id: string;
	    nodePropsList: Record<number, NodeProps>;
	
	    static createFrom(source: any = {}) {
	        return new ResolutionNode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.properties = source["properties"];
	        this.id = source["id"];
	        this.nodePropsList = this.convertValues(source["nodePropsList"], NodeProps, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class StepResult {
	    nodeArray: ResolutionNode[];
	    markedNodes: string[];
	    currentNode: string;
	    finished: boolean;
	
	    static createFrom(source: any = {}) {
	        return new StepResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nodeArray = this.convertValues(source["nodeArray"], ResolutionNode);
	        this.markedNodes = source["markedNodes"];
	        this.currentNode = source["currentNode"];
	        this.finished = source["finished"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

