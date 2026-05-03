export namespace simulation {
	
	export class CommandRequest {
	    commandType: string;
	    targetId: string;
	
	    static createFrom(source: any = {}) {
	        return new CommandRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.commandType = source["commandType"];
	        this.targetId = source["targetId"];
	    }
	}

}

