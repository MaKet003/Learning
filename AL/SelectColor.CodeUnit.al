codeunit 50102 SelectColor
{
    trigger OnRun()
    begin

    end;

    var
        myInt: Integer;

    procedure SelectColor()
    var
        Color: Option Green,Red,Yellow;
    begin
        Color := Color::Red;
        Message('The selected color is: %1', Color);
    end;
}