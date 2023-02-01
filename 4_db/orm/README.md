# 优点：
- 提高开发效率
  - 支持事务、嵌套事务
  - 支持批量的删除
  - 自定义 Logger
  - 比较稳定
  - SQL 构建器，Upsert，数据库锁，Optimizer/Index/Comment Hint，命名参数，子查询
  - hook机制(Before/After Create/Save/Update/Delete/Find)
  - 对象关系Has One, Has Many, Belongs To, Many To Many, Polymorphism
  - 热加载
  - 支持原生sql操作
  - 链式api
# 缺点:
- 牺牲执行性能
- 牺牲灵活性
- 弱化SQL能力