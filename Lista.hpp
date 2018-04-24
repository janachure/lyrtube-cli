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
    if (_ultimo==NULL){
        _ultimo=nuevo;
        nuevo->_valor=elem;
        nuevo->_siguiente=_primero;
        nuevo->_anterior=NULL;
    }else {
        nuevo->_valor = elem;
        nuevo->_siguiente=_primero;
        nuevo->_anterior=NULL;
        _primero->_anterior=nuevo;
        _primero=nuevo;
    }
};

template <typename T>
void Lista<T>::agregarAtras(const T& elem) {
    // Completar
    Nodo* nuevo= new Nodo;
    if (_primero==NULL){
        _primero=nuevo;
    }
    (*nuevo)._valor=elem;
    (*nuevo)._anterior=_ultimo;
    _ultimo=nuevo;
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
    Nodo* nuevo= new Nodo;
    nuevo=_primero;
    T resultado;
    int contador=0;
    while (nuevo->_siguiente!= NULL) {
        if (contador==i) {
            resultado = nuevo->_valor;
            nuevo->_siguiente=NULL;
        };
        nuevo = nuevo-> _siguiente;
        contador++;
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
    Nodo* nuevo= new Nodo;
    nuevo=_primero;
    T resultado;
    int contador=0;
    while (nuevo->_siguiente!= NULL) {
        if (contador==i) {
            resultado = nuevo->_valor;
            nuevo->_siguiente=NULL;
        };
        nuevo = nuevo-> _siguiente;
        contador++;
    }
    return resultado;
}

template <typename T>
void Lista<T>::mostrar(std::ostream& o) {
    // Completar
}
#include "Lista.h"
