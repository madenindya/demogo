import { IProduct } from "./interfaces";

export default class Shirt implements IProduct {

    private name : string;
    private qty : number;
    private color : string;

    constructor(
        name : string,
        qty : number,
    ) {
        this.name = name;
        this.qty = qty;
    }

    public stocks() : number {
        return this.qty;
    }

   public isAvailable() : boolean {
       return this.qty > 0;
   }
}
