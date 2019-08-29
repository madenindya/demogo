import { IProduct } from "./interfaces";
import Book from "./book";
import Shirt from "./shirt";

export default class Main {

    public checkQty(p : IProduct) {
        console.log("Amount: " + p.stocks() + " left!")
    }

    public main() {
        const item1 : IProduct = new Book("buku", 1);
        const item2 : IProduct = new Shirt("baju", 1);

        this.checkQty(item1);
        this.checkQty(item2);

    }


}
