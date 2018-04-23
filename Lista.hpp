#include "Lista.h"

template <typename T>
Lista<T>::Lista() {
    _ultimo=NULL;
    _primero=NULL;
    // Completar
}

// Inicializa una lista vacía y luego utiliza operator= para no duplicar el
// código de la copia de una lista.
template <typename T>
Lista<T>::Lista(const Lista<T>& l) : Lista() { *this = l; }

template <typename T>
Lista<T>::~Lista() {
    // Completar
    while (_primero != NULL) {
        Nodo *p = _primero;
        Nodo *o= _ultimo;
        _primero = _primero->_siguiente;
        _ultimo = _ultimo ->_anterior;
        delete p;
    }
};

template <typename T>
Lista<T>& Lista<T>::operator=(const Lista<T>& l) {
    // Completar

}

template <typename T>
void Lista<T>::agregarAdelante(const T& elem) {
    // Completar
    Nodo* nuevo= new Nodo;
    if (_ultimo==NULL){   // COMO SI RECIEN LO ACABASTE DE CREAR
        _ultimo=nuevo;
    };
    nuevo->_valor=elem;     // ASIGNA EL ELEMENTO PRIMERO
    (*nuevo)._siguiente=_primero;   // EL SIGUIENTE CONECTALO CON EL PRIMERO
    (*nuevo)._anterior=NULL;
    _primero=nuevo;
    //deberia andar
};

template <typename T>
void Lista<T>::agregarAtras(const T& elem) {
    // Completar
    Nodo* nuevo= new Nodo;
    if (_primero==NULL){

    }
    (*nuevo)._valor=elem;
    (*nuevo)._anterior=_ultimo;
    _ultimo=(*nuevo)._anterior;
    (*nuevo)._siguiente=NULL;
};
template <typename T>
int Lista<T>::longitud() const {
    int tam = 0;
    Nodo* p= new Nodo;
    p= _primero;
    while(  p->_siguiente!=NULL){
        p = p -> _siguiente;
        tam++;
    }
    return tam;
};

template <typename T>
const T& Lista<T>::iesimo(Nat i) const {
    // Completar
    Nodo* senhala=_primero;
    T resultado;
    int iterador=0;
    while (iterador!=longitud()+1){
        iterador++;
        resultado=senhala->_valor;
        senhala= _primero->_anterior;
    }
    return resultado;
};

template <typename T>
void Lista<T>::eliminar(Nat i) {
    // Completar
}

template <typename T>
T& Lista<T>::iesimo(Nat i) {
    // Completar (hint: es igual a la anterior...)
}

template <typename T>
void Lista<T>::mostrar(std::ostream& o) {
    // Completar
}
#include "Lista.h"
