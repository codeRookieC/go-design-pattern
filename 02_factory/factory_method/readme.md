### 工厂方法小结

相较于简单工厂模式 工厂方法解决了在扩展组件类时候不符合开闭原则的问题 但是也有相对的缺陷

- 对于每一个组件类 都需要创造对应的工厂类 代码冗余度较高
- 在简单工厂中 在构造多个组件的时候会产生公共切面 我们可以在这里写一些公共逻辑 而在工厂方法中 对应的公共切面不存在 我们只能重复在实现工厂类的时候去写这些逻辑 重复声明