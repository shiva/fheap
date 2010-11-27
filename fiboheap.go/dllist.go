package ring

type Ring struct {
    next, prev *Ring
    Value      interface{} // for use by client; untouched by this library
}


func (r *Ring) init() *Ring {
    r.next = r
    r.prev = r
    return r
}


// Next returns the next ring element. r must not be empty.
func (r *Ring) Next() *Ring {
    if r.next == nil {
        return r.init()
    }
    return r.next
}


