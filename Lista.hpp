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
    if (_primero==NULL){
        _ultimo=nuevo;
        nuevo->_valor=elem;
        nuevo->_siguiente=_primero;
        _primero=nuevo;
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
    if (_primero==NULL)
        return 0;
    while(p!=NULL or _ultimo==NULL){
        tam++;
        p = p -> _siguiente;
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
    while (nuevo!= NULL) {
        if (contador==i) {
            resultado = nuevo->_valor;
            nuevo->_siguiente=NULL;
        };
        nuevo = nuevo-> _siguiente;
        contador++;
