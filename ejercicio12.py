# -*- coding: utf-8 -*-
HORAS_TRABAJADAS_SEMANALES = 48 
nombresEmpleado = input("Ingrese los nombres del empleado: ")
VALOR_HORA = 5000
RETENCION = 12.5/100

# Calcular salario bruto
salarioBruto = HORAS_TRABAJADAS_SEMANALES * VALOR_HORA

# Calcular retención en la fuente
retencion = salarioBruto *  (RETENCION/ 100)

# Calcular salario neto
salarioNeto = salarioBruto - retencion

# Mostrar resultados
print("Nombres:", nombresEmpleado)
print("Salario Bruto:", salarioBruto)
print("Retención en la fuente:", retencion)
print("Salario neto", salarioNeto)