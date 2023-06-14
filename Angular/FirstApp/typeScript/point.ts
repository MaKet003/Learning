
export class Point {
   //() means no Parameters
   constructor(private _x?: number, private _y?: number){
   }

   draw(){
      console.log('X: ' + this._x + ', Y: ' + this._y)
   }

}

