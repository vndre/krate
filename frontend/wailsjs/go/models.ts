export namespace collection_model {
	
	export class Collection {
	    Id: number;
	    Name: string;
	    Genre: string;
	    ImageUrl: string;
	    Progress: number;
	
	    static createFrom(source: any = {}) {
	        return new Collection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Name = source["Name"];
	        this.Genre = source["Genre"];
	        this.ImageUrl = source["ImageUrl"];
	        this.Progress = source["Progress"];
	    }
	}

}

