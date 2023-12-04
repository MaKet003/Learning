codeunit 50100 MyCodeunit
{
    trigger OnRun()
    begin

    end;

    var
        myInt: Integer;

    procedure SelectColor()
    var
        MyFavoriteColor: Enum Color;
        CustomerNames: List of [Text];
        A, B, C: Integer;
    begin
        MyFavoriteColor := Color::Red;
        Message('The selected color is: %1', MyFavoriteColor);

        CustomerNames.Add('Paula');

        C := A + B;
    end;
}