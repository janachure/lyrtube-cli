#include "Lista.h"

template <typename T>
Lista<T>::Lista() {
    _ultimo=NULL;
    _primero=NULL;
    // Completar
}

// Inicializa una lista vacÃ­a y luego utiliza operator= para no duplicar el
// cÃ³digo de la copia de una lista.
template <typename T>
Lista<T>::Lista(const Lista<T>& l) : Lista() { *this = l; }

template <typename T>
Lista<T>::~Lista() {
    // Completar
/*    while (_primero != NULL) {
        Nodo *p = _primero;
        Nodo *o= _ultimo;
        _primero = _primero->_siguiente;
        _ultimo = _ultimo ->_anterior;
        delete p;
    }*/
};

template <typename T>
Lista<T>& Lista<T>::operator=(const Lista<T>& l) {
    return *this;
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
    Nodo *nuevo = new Nodo;
    if (_primero == NULL) {
        _primero = nuevo;
        _ultimo = nuevo;
        nuevo->_valor=elem;
        nuevo->_siguiente = NULL;
        nuevo->_anterior = NULL;
    } else {
        nuevo->_valor = elem;
        _ultimo->_siguiente=nuevo;
        nuevo->_anterior = _ultimo;
        nuevo->_siguiente = NULL;
        _ultimo = nuevo;
    }
};
template <typename T>
int Lista<T>::longitud() const {
    int tam = 0;
    Nodo* p=_primero;
    if (_primero==NULL) {
        return 0;
    }
    while(p!=NULL or _ultimo==NULL){
        tam++;
        p = p -> _siguiente;
    }
    return tam;
};

template <typename T>
const T& Lista<T>::iesimo(Nat i) const {
    // Completar
    Nodo* nuevo=_primero;
    T resultado;
    //int contador=0;
    for (int j = 0; j < longitud(); ++j) {
        if (j==i) {
            j=longitud();
            return nuevo->_valor;
        }
        nuevo = nuevo->_siguiente;
    }
}

template <typename T>
T& Lista<T>::iesimo(Nat i) {
    // Completar

    Nodo* nuevo=_primero;
    for (int j = 0; j < longitud(); ++j) {
        if (j==i) {
            j=longitud();
            return nuevo-> _valor;
        }
        nuevo = nuevo->_siguiente;
    }
};
template <typename T>
void Lista<T>::eliminar(Nat i) {
    Nodo *anterior;
    Nodo *siguiente;
    Nodo *nuevo = _primero;
    if (i == 0) {
        _primero = _primero->_siguiente;
        _ultimo = _primero;
        delete nuevo;
    }
    if (i == longitud()) {
        _ultimo = nuevo->_anterior;
        nuevo->_anterior->_siguiente = NULL;
        delete nuevo;
    } else {
        for (int j = 0; j < longitud() and nuevo->_siguiente!= NULL ; j++) {
            if (i == j) {
                anterior = nuevo->_anterior;
                siguiente = nuevo->_siguiente;
                nuevo->_siguiente->_anterior = anterior;
                nuevo->_anterior->_siguiente = siguiente;
                delete nuevo;
            }
            if(nuevo->_siguiente==NULL)
                break;
            nuevo = nuevo->_siguiente;
        }
    }
};


template <typename T>
void Lista<T>::mostrar(std::ostream& o) {
    // Completar
}
#include "Lista.h"
