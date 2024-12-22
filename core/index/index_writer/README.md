# Java version

```xml
<dependencies>
    <dependency>
        <groupId>org.apache.lucene</groupId>
        <artifactId>lucene-core</artifactId>
        <version>8.11.4</version>
    </dependency>
    <dependency>
        <groupId>org.apache.lucene</groupId>
        <artifactId>lucene-codecs</artifactId>
        <version>8.11.4</version>
    </dependency>
</dependencies>
```


````java
package org.example;

import org.apache.lucene.analysis.Analyzer;
import org.apache.lucene.analysis.standard.StandardAnalyzer;
import org.apache.lucene.codecs.Codec;
import org.apache.lucene.document.Document;
import org.apache.lucene.document.Field;
import org.apache.lucene.document.TextField;
import org.apache.lucene.index.IndexWriter;
import org.apache.lucene.index.IndexWriterConfig;
import org.apache.lucene.search.similarities.BM25Similarity;
import org.apache.lucene.store.Directory;
import org.apache.lucene.store.FSDirectory;
import org.apache.lucene.codecs.simpletext.SimpleTextCodec;

import java.nio.file.Paths;

public class Writer {
    public static void main(String[] args) throws Exception {
        String indexPath = "data";
        Directory dir = FSDirectory.open(Paths.get(indexPath));

        Analyzer analyzer = new StandardAnalyzer();

        IndexWriterConfig iwc = new IndexWriterConfig(analyzer);
        iwc.setOpenMode(IndexWriterConfig.OpenMode.CREATE);

        Codec codec = new SimpleTextCodec();
        iwc.setCodec(codec);

        BM25Similarity similarity = new BM25Similarity();
        iwc.setSimilarity(similarity);

        IndexWriter indexWriter = new IndexWriter(dir, iwc);

        Document doc;

        doc = new Document();
        doc.add(new TextField("a", "74", Field.Store.YES));
        doc.add(new TextField("a1", "86", Field.Store.YES));
        doc.add(new TextField("a2", "1237", Field.Store.YES));
        indexWriter.addDocument(doc);
        // 1
        doc = new Document();
        doc.add(new TextField("a", "123", Field.Store.YES));
        doc.add(new TextField("a1", "123", Field.Store.YES));
        doc.add(new TextField("a2", "789", Field.Store.YES));
        indexWriter.addDocument(doc);
        // 2
        doc = new Document();
        doc.add(new TextField("a", "741", Field.Store.YES));
        doc.add(new TextField("a1", "861", Field.Store.YES));
        doc.add(new TextField("a2", "12137", Field.Store.YES));
        indexWriter.addDocument(doc);

        indexWriter.commit();
        indexWriter.close();
    }
}
````