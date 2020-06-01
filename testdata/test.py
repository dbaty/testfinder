class TestEmptyClass:
    pass


class TestClassWithMethods:

    def test_method1(self):
        pass

    def test_method2(self):
        def test_inner_function_to_ignore():
            pass


def test_func():
    pass
