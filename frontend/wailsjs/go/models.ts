export namespace recipe {
	
	export class Ingredient {
	    numerator: number;
	    denominator: number;
	    unit: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Ingredient(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.numerator = source["numerator"];
	        this.denominator = source["denominator"];
	        this.unit = source["unit"];
	        this.name = source["name"];
	    }
	}
	export class Recipe {
	    title: string;
	    author: string;
	    ingredients: Ingredient[];
	    directions: string[];
	    tags: {[key: string]: boolean};
	    id: number;
	
	    static createFrom(source: any = {}) {
	        return new Recipe(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.author = source["author"];
	        this.ingredients = this.convertValues(source["ingredients"], Ingredient);
	        this.directions = source["directions"];
	        this.tags = source["tags"];
	        this.id = source["id"];
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

