#ifndef _MY_PACKAGE_FOO_HPP_
#define _MY_PACKAGE_FOO_HPP_


namespace rozen{


	class cxxFoo {
	public:
		int a;
		cxxFoo(int _a):a(_a){};
		~cxxFoo(){};
		void Bar();
	};
}

#endif
