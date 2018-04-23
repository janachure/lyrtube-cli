#ifndef _LISTA_ALUMNOS_H_
#define _LISTA_ALUMNOS_H_

#include <ostream>

typedef unsigned long Nat;

template <typename T>
class Lista {
public:

    /* Constructor por defecto de la clase Lista. */
    Lista();

    /* Constructor por copia de la clase Lista. */
    Lista(const Lista& l);

    /* Destructor de la clase Lista. */
    ~Lista();

    /* Operador asignación */
    Lista& operator=(const Lista& aCopiar);

    /* Agrega un elemento al principio de la lista. */
    void agregarAdelante(const T& elem);

    /* Agrega un elemento al final de la lista. */
    void agregarAtras(const T& elem);

    /* Devuelve la cantidad de elementos que contiene la lista. */
    int longitud() const;

    /* Devuelve el elemento en la i-ésima posición de la lista,
     * como una referencia inmutable.
     * Precondición: 0 <= i < this->longitud */
    const T& iesimo(Nat i) const;

    /* Devuelve el elemento en la i-ésima posición de la lista,
     * como una referencia mutable.
     * Precondición: 0 <= i < this->longitud */
    T& iesimo(Nat i);

    /* Elimina el i-ésimo elemento de la Lista.
     * Precondición: 0 <= i < this->longitud */
    void eliminar(Nat i);

    /* Muestra la lista en un std::ostream
     * formato de salida: [a_0, a_1, a_2, ...]
     * donde a_i es el resultado del "mostrar" del alumno i de la lista (ya definido)
     * @param o
     */
    void mostrar(std::ostream& o);

    /**
     * Utiliza el método mostrar(os) para sobrecargar el operador <<
     */
    friend std::ostream& operator<<(std::ostream& os, Lista& l) {
        l.mostrar(os);
        return os;
    }

private:

    struct Nodo {
        Nodo* _siguiente;
        T _valor;
        Nodo* _anterior;
    };
    Nodo* _primero;
    Nodo* _ultimo;
    // Completar
};

#include "Lista.hpp"

#endif
