# population_gen_go
A random program that generates a virtual population data given initial parameters.
Uses fyne and excelize library.

# How to use:
- Set initial parameters using the GUI.
- Program then outputs the results as an xlsx file.
-   The workbook will contain age, offsprings, born year, death year, and other stats.

!! Program may run longer depending on the initial parameters. The program only finishes as long as the population being simulated continues to live and stops when it meets the timer. It will prematurely end if the population did not survive even when the timer parameter is set.

!! Depending how big the population gets, the excel file may not hold all of the population individuals, or program may crash if excel row is reached.


# GUI
![GUI](https://user-images.githubusercontent.com/93850550/165607854-2e65cdd5-30a9-4f7f-955c-5333b4e87ec6.png)

# Data Sample
![image](https://user-images.githubusercontent.com/93850550/165641765-d4221b79-046d-4fe9-8e2e-bf003f228aeb.png)

