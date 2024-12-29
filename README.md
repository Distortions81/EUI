## A Minimal vector-based Ebitenengine retained-mode UI library
### This library is no where near complete (PRE-ALPHA), and breaking changes will come very often.
*For quicker development, the code is not composed as a library at this time*  
Drawn using vectors and supports floating-point UI scaling.

![demo animation](https://github.com/user-attachments/assets/eef712c8-fa1e-4afe-826e-624c860ab842)

# Windows, and flows:
window -> main-flow -> sub-flow  
Each flow has a direction, horizontal or vertical.

![flows-screenshot](https://github.com/user-attachments/assets/dcc79179-361d-420c-959f-c1785433bb5b)



# Flow and item pinning:
Flows and items can have pinning type set. The default is PIN_TOP_LEFT.  
The pin changes the point of refrence for the item/flow position.

Other options are:  
PIN_TOP_LEFT  
PIN_TOP_CENTER  
PIN_TOP_RIGHT

PIN_MID_LEFT  
PIN_MID_CENTER  
PIN_MID_RIGHT

PIN_BOTTOM_LEFT  
PIN_BOTTOM_CENTER  
PIN_BOTTOM_RIGHT
