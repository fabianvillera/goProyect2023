# -*- coding: utf-8 -*-
# Solicitamos al usuario que ingrese la edad de Juan
juan = int(input("Ingresa la edad de Juan: "))

# Calculamos las edades de Alberto y Ana según las proporciones dadas
alberto = (2 * juan) // 3
ana = (4 * juan) // 3

# Calculamos la edad de la mamá como la suma de todas las edades
mama = juan + alberto + ana

# Verificamos si la mamá tiene tres hijos
if mama == 3 * juan:
    print("Edad de Alberto:", alberto)
    print("Edad de Ana:", ana)
    print("Edad de la mamá:", mama)
else:
    print("No se encontraron edades que cumplan con las condiciones.")
