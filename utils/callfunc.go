package utils

import (
    "errors"
    "log"
    "reflect"
    "strconv"
)

func GetNumIn(function interface{}) (int, error) {
    value := reflect.ValueOf(function)
    if value.Kind() != reflect.Func {
        return -1, errors.New("not a function")
    }
    return value.Type().NumIn(), nil
}

func GetNumOut(function interface{}) (int, error) {
    value := reflect.ValueOf(function)
    if value.Kind() != reflect.Func {
        return -1, errors.New("not a function")
    }
    return value.Type().NumOut(), nil
}

func Call(function interface{}, args ...string) (output []string) {
    value := reflect.ValueOf(function)
    if value.Kind() != reflect.Func {
        log.Println("function is not function")
        return
    }

    parameters := make([]reflect.Type, 0, value.Type().NumIn())
    for i := 0; i < value.Type().NumIn(); i++ {
        arg := value.Type().In(i)
        //log.Printf("argument %d is %s[%s] type \n", i, arg.Kind(), arg.Name())
        parameters = append(parameters, arg)
    }

    if value.Type().NumIn() != len(args) {
        log.Printf("argument %d length doesn't equal to provide length %d \n", value.Type().NumIn(), len(args))
        return
    }

    outs := make([]reflect.Type, 0, value.Type().NumOut())
    for i := 0; i < value.Type().NumOut(); i++ {
        arg := value.Type().Out(i)
        //log.Printf("out %d is %s[%s] type \n", i, arg.Kind(), arg.Name())
        outs = append(outs, arg)
    }

    if value.Type().NumOut() < 1 {
        log.Println("outs length must greater than 0")
        return
    }

    argValues := make([]reflect.Value, 0, len(parameters))
    for i := 0; i < len(args); i++ {
        switch parameters[i] {
        case reflect.TypeOf(int(0)):
            v, err := strconv.ParseInt(args[i], 10, 64)
            if err != nil {
                log.Printf("argument %d %s convert int failed, %v \n", i, args[i], err)
                return
            }
            argValues = append(argValues, reflect.ValueOf(int(v)))
        case reflect.TypeOf(string("")):
            argValues = append(argValues, reflect.ValueOf(args[i]))
        default:
            log.Printf("unsupport type %s[%s] \n", parameters[i].Kind(), parameters[i].Name())
            return
        }
    }

    resultValues := value.Call(argValues)
    
    for i := 0; i < len(resultValues); i++ {
        switch resultValues[i].Type() {
            case reflect.TypeOf(int(0)):
                //log.Println("result: ", i, ", ", resultValues[i].Interface().(int))
                output = append(output, strconv.Itoa(resultValues[i].Interface().(int)))
            case reflect.TypeOf(string("")):
                //log.Println("result: ", i, ", ", resultValues[i].Interface().(string))
                output = append(output, resultValues[i].Interface().(string))
            default:
                log.Printf("type: %s[%s], value: %v \n", resultValues[i].Type().Kind(), resultValues[i].Type().Name(), resultValues[i].Interface())
            }
    }
    return
}